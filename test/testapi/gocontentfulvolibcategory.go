// Code generated by https://github.com/foomo/gocontentful 1.0.20 - DO NOT EDIT.
package testapi

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/foomo/contentful"
)

const ContentTypeCategory = "category"

// ---Category private methods---

// ---Category public methods---

func (cc *ContentfulClient) GetAllCategory() (voMap map[string]*CfCategory, err error) {
	if cc == nil {
		return nil, errors.New("GetAllCategory: No client available")
	}
	cc.cacheMutex.sharedDataGcLock.RLock()
	cacheInit := cc.cacheInit
	optimisticPageSize := cc.optimisticPageSize
	cc.cacheMutex.sharedDataGcLock.RUnlock()
	if cacheInit {
		return cc.Cache.entryMaps.category, nil
	}
	col, err := cc.optimisticPageSizeGetAll("category", optimisticPageSize)
	if err != nil {
		return nil, err
	}
	allCategory, err := colToCfCategory(col, cc)
	if err != nil {
		return nil, err
	}
	categoryMap := map[string]*CfCategory{}
	for _, category := range allCategory {
		categoryMap[category.Sys.ID] = category
	}
	return categoryMap, nil
}

func (cc *ContentfulClient) GetFilteredCategory(query *contentful.Query) (voMap map[string]*CfCategory, err error) {
	if cc == nil || cc.Client == nil {
		return nil, errors.New("getFilteredCategory: No client available")
	}
	col := cc.Client.Entries.List(cc.SpaceID)
	if query != nil {
		col.Query = *query
	}
	col.Query.ContentType("category").Locale("*").Include(0)
	_, err = col.GetAll()
	if err != nil {
		return nil, errors.New("getFilteredCategory: " + err.Error())
	}
	allCategory, err := colToCfCategory(col, cc)
	if err != nil {
		return nil, errors.New("getFilteredCategory: " + err.Error())
	}
	categoryMap := map[string]*CfCategory{}
	for _, category := range allCategory {
		categoryMap[category.Sys.ID] = category
	}
	return categoryMap, nil
}

func (cc *ContentfulClient) GetCategoryByID(id string, forceNoCache ...bool) (vo *CfCategory, err error) {
	if cc == nil || cc.Client == nil {
		return nil, errors.New("GetCategoryByID: No client available")
	}
	if cc.cacheInit && (len(forceNoCache) == 0 || !forceNoCache[0]) {
		cc.cacheMutex.categoryGcLock.RLock()
		defer cc.cacheMutex.categoryGcLock.RUnlock()
		vo, ok := cc.Cache.entryMaps.category[id]
		if ok {
			return vo, nil
		}
		return nil, fmt.Errorf("GetCategoryByID: entry '%s' not found in cache", id)
	}
	col := cc.Client.Entries.List(cc.SpaceID)
	col.Query.ContentType("category").Locale("*").Include(0).Equal("sys.id", id)
	_, err = col.GetAll()
	if err != nil {
		return nil, err
	}
	if len(col.Items) == 0 {
		return nil, fmt.Errorf("GetCategoryByID: %s Not found", id)
	}
	vos, err := colToCfCategory(col, cc)
	if err != nil {
		return nil, fmt.Errorf("GetCategoryByID: Error converting %s to VO: %w", id, err)
	}
	vo = vos[0]
	return
}

func NewCfCategory(contentfulClient ...*ContentfulClient) (cfCategory *CfCategory) {
	cfCategory = &CfCategory{}
	if len(contentfulClient) != 0 && contentfulClient[0] != nil {
		cfCategory.CC = contentfulClient[0]
	}

	cfCategory.Fields.Title = map[string]string{}

	cfCategory.Fields.Icon = map[string]ContentTypeSys{}

	cfCategory.Fields.CategoryDescription = map[string]string{}

	cfCategory.Sys.ContentType.Sys.ID = "category"
	cfCategory.Sys.ContentType.Sys.Type = FieldTypeLink
	cfCategory.Sys.ContentType.Sys.LinkType = "ContentType"
	return
}
func (vo *CfCategory) GetParents(contentType ...string) (parents []EntryReference, err error) {
	if vo == nil {
		return nil, errors.New("GetParents: Value Object is nil")
	}
	if vo.CC == nil {
		return nil, errors.New("GetParents: Value Object has no Contentful Client set")
	}
	return commonGetParents(vo.CC, vo.Sys.ID, contentType)
}

func (vo *CfCategory) GetPublishingStatus() string {
	if vo == nil {
		return ""
	}
	if vo.Sys.PublishedVersion == 0 {
		return StatusDraft
	}
	if vo.Sys.Version-vo.Sys.PublishedVersion == 1 {
		return StatusPublished
	}
	return StatusChanged
}

// Category Field getters

func (vo *CfCategory) Title(locale ...Locale) string {
	if vo == nil {
		return ""
	}
	if vo.CC == nil {
		return ""
	}
	vo.Fields.RWLockTitle.RLock()
	defer vo.Fields.RWLockTitle.RUnlock()
	loc := defaultLocale
	if len(locale) != 0 {
		loc = locale[0]
		if _, ok := localeFallback[loc]; !ok {
			if vo.CC.logFn != nil && vo.CC.logLevel <= LogError {
				vo.CC.logFn(map[string]interface{}{"content type": vo.Sys.ContentType.Sys.ID, "entry ID": vo.Sys.ID, "method": "Title()"}, LogError, ErrLocaleUnsupported)
			}
			return ""
		}
	}
	if _, ok := vo.Fields.Title[string(loc)]; !ok {
		if _, ok := localeFallback[loc]; !ok {
			if vo.CC.logFn != nil && vo.CC.logLevel == LogDebug {
				vo.CC.logFn(map[string]interface{}{"content type": vo.Sys.ContentType.Sys.ID, "entry ID": vo.Sys.ID, "method": "Title()"}, LogWarn, ErrNotSet)
			}
			return ""
		}
		loc = localeFallback[loc]
		if _, ok := vo.Fields.Title[string(loc)]; !ok {
			if vo.CC.logFn != nil && vo.CC.logLevel == LogDebug {
				vo.CC.logFn(map[string]interface{}{"content type": vo.Sys.ContentType.Sys.ID, "entry ID": vo.Sys.ID, "method": "Title()"}, LogWarn, ErrNotSetNoFallback)
			}
			return ""
		}
	}
	return vo.Fields.Title[string(loc)]
}

func (vo *CfCategory) Icon(locale ...Locale) *contentful.AssetNoLocale {
	if vo == nil {
		return nil
	}
	if vo.CC == nil {
		return nil
	}
	vo.Fields.RWLockIcon.RLock()
	defer vo.Fields.RWLockIcon.RUnlock()
	loc := defaultLocale
	reqLoc := defaultLocale
	if len(locale) != 0 {
		loc = locale[0]
		reqLoc = locale[0]
		if _, ok := localeFallback[loc]; !ok {
			if vo.CC.logFn != nil && vo.CC.logLevel <= LogError {
				vo.CC.logFn(map[string]interface{}{"content type": vo.Sys.ContentType.Sys.ID, "entry ID": vo.Sys.ID, "method": "Icon()"}, LogError, ErrLocaleUnsupported)
			}
			return nil
		}
	}
	if _, ok := vo.Fields.Icon[string(loc)]; !ok {
		if _, ok := localeFallback[loc]; !ok {
			if vo.CC.logFn != nil && vo.CC.logLevel == LogDebug {
				vo.CC.logFn(map[string]interface{}{"content type": vo.Sys.ContentType.Sys.ID, "entry ID": vo.Sys.ID, "method": "Icon()"}, LogWarn, ErrNotSet)
			}
			return nil
		}
		loc = localeFallback[loc]
		if _, ok := vo.Fields.Icon[string(loc)]; !ok {
			if vo.CC.logFn != nil && vo.CC.logLevel == LogDebug {
				vo.CC.logFn(map[string]interface{}{"content type": vo.Sys.ContentType.Sys.ID, "entry ID": vo.Sys.ID, "method": "Icon()"}, LogWarn, ErrNotSetNoFallback)
			}
			return nil
		}
	}
	localizedIcon := vo.Fields.Icon[string(loc)]
	asset, err := vo.CC.GetAssetByID(localizedIcon.Sys.ID)
	if err != nil {
		if vo.CC.logFn != nil && vo.CC.logLevel == LogDebug {
			vo.CC.logFn(map[string]interface{}{"content type": vo.Sys.ContentType.Sys.ID, "entry ID": vo.Sys.ID, "method": "Icon()"}, LogError, ErrNoTypeOfRefAsset)
		}
		return nil
	}
	tempAsset := &contentful.AssetNoLocale{}
	tempAsset.Sys = asset.Sys
	tempAsset.Fields = &contentful.FileFieldsNoLocale{}
	if _, ok := asset.Fields.Title[string(reqLoc)]; ok {
		tempAsset.Fields.Title = asset.Fields.Title[string(reqLoc)]
	} else {
		tempAsset.Fields.Title = asset.Fields.Title[string(loc)]
	}
	if _, ok := asset.Fields.Description[string(reqLoc)]; ok {
		tempAsset.Fields.Description = asset.Fields.Description[string(reqLoc)]
	} else {
		tempAsset.Fields.Description = asset.Fields.Description[string(loc)]
	}
	if _, ok := asset.Fields.File[string(reqLoc)]; ok {
		tempAsset.Fields.File = asset.Fields.File[string(reqLoc)]
	} else {
		tempAsset.Fields.File = asset.Fields.File[string(loc)]
	}
	return tempAsset
}

func (vo *CfCategory) CategoryDescription(locale ...Locale) string {
	if vo == nil {
		return ""
	}
	if vo.CC == nil {
		return ""
	}
	vo.Fields.RWLockCategoryDescription.RLock()
	defer vo.Fields.RWLockCategoryDescription.RUnlock()
	loc := defaultLocale
	if len(locale) != 0 {
		loc = locale[0]
		if _, ok := localeFallback[loc]; !ok {
			if vo.CC.logFn != nil && vo.CC.logLevel <= LogError {
				vo.CC.logFn(map[string]interface{}{"content type": vo.Sys.ContentType.Sys.ID, "entry ID": vo.Sys.ID, "method": "CategoryDescription()"}, LogError, ErrLocaleUnsupported)
			}
			return ""
		}
	}
	if _, ok := vo.Fields.CategoryDescription[string(loc)]; !ok {
		if _, ok := localeFallback[loc]; !ok {
			if vo.CC.logFn != nil && vo.CC.logLevel == LogDebug {
				vo.CC.logFn(map[string]interface{}{"content type": vo.Sys.ContentType.Sys.ID, "entry ID": vo.Sys.ID, "method": "CategoryDescription()"}, LogWarn, ErrNotSet)
			}
			return ""
		}
		loc = localeFallback[loc]
		if _, ok := vo.Fields.CategoryDescription[string(loc)]; !ok {
			if vo.CC.logFn != nil && vo.CC.logLevel == LogDebug {
				vo.CC.logFn(map[string]interface{}{"content type": vo.Sys.ContentType.Sys.ID, "entry ID": vo.Sys.ID, "method": "CategoryDescription()"}, LogWarn, ErrNotSetNoFallback)
			}
			return ""
		}
	}
	return vo.Fields.CategoryDescription[string(loc)]
}

// Category Field setters

func (vo *CfCategory) SetTitle(title string, locale ...Locale) (err error) {
	if vo == nil {
		return errors.New("SetTitle(title: Value Object is nil")
	}
	loc := defaultLocale
	if len(locale) != 0 {
		loc = locale[0]
		if _, ok := localeFallback[loc]; !ok {
			return ErrLocaleUnsupported
		}
	}
	vo.Fields.RWLockTitle.Lock()
	defer vo.Fields.RWLockTitle.Unlock()
	if vo.Fields.Title == nil {
		vo.Fields.Title = make(map[string]string)
	}
	vo.Fields.Title[string(loc)] = title
	return
}

func (vo *CfCategory) SetIcon(icon ContentTypeSys, locale ...Locale) (err error) {
	if vo == nil {
		return errors.New("SetIcon(icon: Value Object is nil")
	}
	loc := defaultLocale
	if len(locale) != 0 {
		loc = locale[0]
		if _, ok := localeFallback[loc]; !ok {
			return ErrLocaleUnsupported
		}
	}
	vo.Fields.RWLockIcon.Lock()
	defer vo.Fields.RWLockIcon.Unlock()
	if vo.Fields.Icon == nil {
		vo.Fields.Icon = make(map[string]ContentTypeSys)
	}
	vo.Fields.Icon[string(loc)] = icon
	return
}

func (vo *CfCategory) SetCategoryDescription(categoryDescription string, locale ...Locale) (err error) {
	if vo == nil {
		return errors.New("SetCategoryDescription(categoryDescription: Value Object is nil")
	}
	loc := defaultLocale
	if len(locale) != 0 {
		loc = locale[0]
		if _, ok := localeFallback[loc]; !ok {
			return ErrLocaleUnsupported
		}
	}
	vo.Fields.RWLockCategoryDescription.Lock()
	defer vo.Fields.RWLockCategoryDescription.Unlock()
	if vo.Fields.CategoryDescription == nil {
		vo.Fields.CategoryDescription = make(map[string]string)
	}
	vo.Fields.CategoryDescription[string(loc)] = categoryDescription
	return
}

func (vo *CfCategory) UpsertEntry() (err error) {
	if vo == nil {
		return errors.New("UpsertEntry: Value Object is nil")
	}
	if vo.CC == nil {
		return errors.New("UpsertEntry: Value Object has nil Contentful client")
	}
	if vo.CC.clientMode != ClientModeCMA {
		return errors.New("UpsertEntry: Only available in ClientModeCMA")
	}
	cfEntry := &contentful.Entry{}
	tmp, errMarshal := json.Marshal(vo)
	if errMarshal != nil {
		return errors.New("CfCategory UpsertEntry: Can't marshal JSON from VO")
	}
	errUnmarshal := json.Unmarshal(tmp, &cfEntry)
	if errUnmarshal != nil {
		return errors.New("CfCategory UpsertEntry: Can't unmarshal JSON into CF entry")
	}

	err = vo.CC.Client.Entries.Upsert(vo.CC.SpaceID, cfEntry)
	if err != nil {
		return fmt.Errorf("CfCategory UpsertEntry: Operation failed: %w", err)
	}
	return
}
func (vo *CfCategory) PublishEntry() (err error) {
	if vo == nil {
		return errors.New("PublishEntry: Value Object is nil")
	}
	if vo.CC == nil {
		return errors.New("PublishEntry: Value Object has nil Contentful client")
	}
	if vo.CC.clientMode != ClientModeCMA {
		return errors.New("PublishEntry: Only available in ClientModeCMA")
	}
	cfEntry := &contentful.Entry{}
	tmp, errMarshal := json.Marshal(vo)
	if errMarshal != nil {
		return errors.New("CfCategory PublishEntry: Can't marshal JSON from VO")
	}
	errUnmarshal := json.Unmarshal(tmp, &cfEntry)
	if errUnmarshal != nil {
		return errors.New("CfCategory PublishEntry: Can't unmarshal JSON into CF entry")
	}
	err = vo.CC.Client.Entries.Publish(vo.CC.SpaceID, cfEntry)
	if err != nil {
		return fmt.Errorf("CfCategory PublishEntry: publish operation failed: %w", err)
	}
	return
}
func (vo *CfCategory) UnpublishEntry() (err error) {
	if vo == nil {
		return errors.New("UnpublishEntry: Value Object is nil")
	}
	if vo.CC == nil {
		return errors.New("UnpublishEntry: Value Object has nil Contentful client")
	}
	if vo.CC.clientMode != ClientModeCMA {
		return errors.New("UnpublishEntry: Only available in ClientModeCMA")
	}
	cfEntry := &contentful.Entry{}
	tmp, errMarshal := json.Marshal(vo)
	if errMarshal != nil {
		return errors.New("CfCategory UnpublishEntry: Can't marshal JSON from VO")
	}
	errUnmarshal := json.Unmarshal(tmp, &cfEntry)
	if errUnmarshal != nil {
		return errors.New("CfCategory UnpublishEntry: Can't unmarshal JSON into CF entry")
	}
	err = vo.CC.Client.Entries.Unpublish(vo.CC.SpaceID, cfEntry)
	if err != nil {
		return fmt.Errorf("CfCategory UnpublishEntry: unpublish operation failed: %w", err)
	}
	return
}
func (vo *CfCategory) UpdateEntry() (err error) {
	if vo == nil {
		return errors.New("UpdateEntry: Value Object is nil")
	}
	if vo.CC == nil {
		return errors.New("UpdateEntry: Value Object has nil Contentful client")
	}
	if vo.CC.clientMode != ClientModeCMA {
		return errors.New("UpdateEntry: Only available in ClientModeCMA")
	}
	cfEntry := &contentful.Entry{}
	tmp, errMarshal := json.Marshal(vo)
	if errMarshal != nil {
		return errors.New("CfCategory UpdateEntry: Can't marshal JSON from VO")
	}
	errUnmarshal := json.Unmarshal(tmp, &cfEntry)
	if errUnmarshal != nil {
		return errors.New("CfCategory UpdateEntry: Can't unmarshal JSON into CF entry")
	}

	err = vo.CC.Client.Entries.Upsert(vo.CC.SpaceID, cfEntry)
	if err != nil {
		return fmt.Errorf("CfCategory UpdateEntry: upsert operation failed: %w", err)
	}
	tmp, errMarshal = json.Marshal(cfEntry)
	if errMarshal != nil {
		return errors.New("CfCategory UpdateEntry: Can't marshal JSON back from CF entry")
	}
	errUnmarshal = json.Unmarshal(tmp, &vo)
	if errUnmarshal != nil {
		return errors.New("CfCategory UpdateEntry: Can't unmarshal JSON back into VO")
	}
	err = vo.CC.Client.Entries.Publish(vo.CC.SpaceID, cfEntry)
	if err != nil {
		return fmt.Errorf("CfCategory UpdateEntry: publish operation failed: %w", err)
	}
	return
}
func (vo *CfCategory) DeleteEntry() (err error) {
	if vo == nil {
		return errors.New("DeleteEntry: Value Object is nil")
	}
	if vo.CC == nil {
		return errors.New("DeleteEntry: Value Object has nil Contentful client")
	}
	if vo.CC.clientMode != ClientModeCMA {
		return errors.New("DeleteEntry: Only available in ClientModeCMA")
	}
	cfEntry := &contentful.Entry{}
	tmp, errMarshal := json.Marshal(vo)
	if errMarshal != nil {
		return errors.New("CfCategory DeleteEntry: Can't marshal JSON from VO")
	}
	errUnmarshal := json.Unmarshal(tmp, &cfEntry)
	if errUnmarshal != nil {
		return errors.New("CfCategory DeleteEntry: Can't unmarshal JSON into CF entry")
	}
	if cfEntry.Sys.PublishedCounter > 0 {
		errUnpublish := vo.CC.Client.Entries.Unpublish(vo.CC.SpaceID, cfEntry)
		if errUnpublish != nil && !strings.Contains(errUnpublish.Error(), "Not published") {
			return fmt.Errorf("CfCategory DeleteEntry: Unpublish entry failed: %w", errUnpublish)
		}
	}
	errDelete := vo.CC.Client.Entries.Delete(vo.CC.SpaceID, cfEntry.Sys.ID)
	if errDelete != nil {
		return fmt.Errorf("CfCategory DeleteEntry: Delete entry failed: %w", errDelete)
	}
	return nil
}
func (vo *CfCategory) ToReference() (refSys ContentTypeSys) {
	if vo == nil {
		return refSys
	}
	refSys.Sys.ID = vo.Sys.ID
	refSys.Sys.Type = FieldTypeLink
	refSys.Sys.LinkType = FieldLinkTypeEntry
	return
}

func (cc *ContentfulClient) cacheAllCategory(ctx context.Context, resultChan chan<- ContentTypeResult) (vos map[string]*CfCategory, err error) {
	if cc == nil || cc.Client == nil {
		return nil, errors.New("cacheAllCategory: No CDA/CPA client available")
	}
	var allCategory []*CfCategory
	col := &contentful.Collection{
		Items: []interface{}{},
	}
	cc.cacheMutex.sharedDataGcLock.RLock()
	defer cc.cacheMutex.sharedDataGcLock.RUnlock()
	if cc.offline {
		for _, entry := range cc.offlineTemp.Entries {
			if entry.Sys.ContentType.Sys.ID == ContentTypeCategory {
				col.Items = append(col.Items, entry)
			}
		}
	} else {
		col, err = cc.optimisticPageSizeGetAll("category", cc.optimisticPageSize)
		if err != nil {
			return nil, errors.New("optimisticPageSizeGetAll for Category failed: " + err.Error())
		}
	}
	allCategory, err = colToCfCategory(col, cc)
	if err != nil {
		return nil, errors.New("colToCfCategory failed: " + err.Error())
	}
	categoryMap := map[string]*CfCategory{}
	for _, category := range allCategory {
		if cc.cacheInit {
			existingCategory, err := cc.GetCategoryByID(category.Sys.ID)
			if err == nil && existingCategory != nil && existingCategory.Sys.Version > category.Sys.Version {
				return nil, fmt.Errorf("cache update canceled because Category entry %s is newer in cache", category.Sys.ID)
			}
		}
		categoryMap[category.Sys.ID] = category
		result := ContentTypeResult{
			EntryID:     category.Sys.ID,
			ContentType: ContentTypeCategory,
			References:  map[string][]EntryReference{},
		}
		addEntry := func(id string, refs EntryReference) {
			if result.References[id] == nil {
				result.References[id] = []EntryReference{}
			}
			result.References[id] = append(result.References[id], refs)
		}
		_ = addEntry

		if ctx.Err() != nil {
			return nil, ctx.Err()
		}
		resultChan <- result
	}
	return categoryMap, nil
}

func (cc *ContentfulClient) cacheCategoryByID(ctx context.Context, id string, entryPayload *contentful.Entry, entryDelete bool) error {
	cc.cacheMutex.categoryGcLock.Lock()
	defer cc.cacheMutex.categoryGcLock.Unlock()
	cc.cacheMutex.idContentTypeMapGcLock.Lock()
	defer cc.cacheMutex.idContentTypeMapGcLock.Unlock()
	cc.cacheMutex.parentMapGcLock.Lock()
	defer cc.cacheMutex.parentMapGcLock.Unlock()

	var col *contentful.Collection
	if entryPayload != nil {
		col = &contentful.Collection{
			Items: []interface{}{entryPayload},
		}
		id = entryPayload.Sys.ID
	} else {
		if cc.Client == nil {
			return errors.New("cacheCategoryByID: No client available")
		}
		if !entryDelete {
			col = cc.Client.Entries.List(cc.SpaceID)
			col.Query.ContentType("category").Locale("*").Include(0).Equal("sys.id", id)
			_, err := col.GetAll()
			if err != nil {
				return err
			}
		}
	}
	// It was deleted
	if col != nil && len(col.Items) == 0 || entryDelete {
		delete(cc.Cache.entryMaps.category, id)
		delete(cc.Cache.idContentTypeMap, id)
		// delete as child
		delete(cc.Cache.parentMap, id)
		// delete as parent
		for childID, parents := range cc.Cache.parentMap {
			newParents := []EntryReference{}
			for _, parent := range parents {
				if parent.ID != id {
					newParents = append(newParents, parent)
				}
			}
			cc.Cache.parentMap[childID] = newParents
		}
		return nil
	}
	vos, err := colToCfCategory(col, cc)
	if err != nil {
		return fmt.Errorf("cacheCategoryByID: Error converting %s to VO: %w", id, err)
	}
	category := vos[0]
	if cc.Cache.entryMaps.category == nil {
		cc.Cache.entryMaps.category = map[string]*CfCategory{}
	}
	cc.Cache.entryMaps.category[id] = category
	cc.Cache.idContentTypeMap[id] = category.Sys.ContentType.Sys.ID
	allChildrensIds := map[string]bool{}

	_ = allChildrensIds // safety net
	// clean up child-parents that don't exist anymore
	for childID, parents := range cc.Cache.parentMap {
		if _, isCollectedChildID := allChildrensIds[childID]; isCollectedChildID {
			continue
		}
		newParents := []EntryReference{}
		for _, parent := range parents {
			if parent.ID != id {
				newParents = append(newParents, parent)
			}
		}
		cc.Cache.parentMap[childID] = newParents
	}
	return nil
}

func colToCfCategory(col *contentful.Collection, cc *ContentfulClient) (vos []*CfCategory, err error) {
	for _, item := range col.Items {
		var vo CfCategory
		byteArray, _ := json.Marshal(item)
		err = json.NewDecoder(bytes.NewReader(byteArray)).Decode(&vo)
		if err != nil {
			break
		}
		if cc.textJanitor {

			vo.Fields.Title = cleanUpStringField(vo.Fields.Title)

			vo.Fields.CategoryDescription = cleanUpStringField(vo.Fields.CategoryDescription)

		}
		vo.CC = cc
		vos = append(vos, &vo)
	}
	return vos, err
}
