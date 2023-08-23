"use strict";
/* tslint:disable */
/* eslint-disable */
/**
 * OPERA Cloud CRM Configuration API
 * APIs for Customer Relationship Management (profile) configuration, such as creating preferences, or address types.  It also includes Membership Configuration, where you can retrieve membership levels that are configured for a property, or create new membership enrollment codes.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 21.5.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
Object.defineProperty(exports, "__esModule", { value: true });
exports.CityPostalCodesDetailsToJSON = exports.CityPostalCodesDetailsFromJSONTyped = exports.CityPostalCodesDetailsFromJSON = exports.instanceOfCityPostalCodesDetails = void 0;
const runtime_1 = require("../runtime");
const CityPostalCodeType_1 = require("./CityPostalCodeType");
const InstanceLink_1 = require("./InstanceLink");
const MasterInfoType_1 = require("./MasterInfoType");
const WarningType_1 = require("./WarningType");
/**
 * Check if a given object implements the CityPostalCodesDetails interface.
 */
function instanceOfCityPostalCodesDetails(value) {
    let isInstance = true;
    return isInstance;
}
exports.instanceOfCityPostalCodesDetails = instanceOfCityPostalCodesDetails;
function CityPostalCodesDetailsFromJSON(json) {
    return CityPostalCodesDetailsFromJSONTyped(json, false);
}
exports.CityPostalCodesDetailsFromJSON = CityPostalCodesDetailsFromJSON;
function CityPostalCodesDetailsFromJSONTyped(json, ignoreDiscriminator) {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        'cityPostalCodes': !(0, runtime_1.exists)(json, 'cityPostalCodes') ? undefined : (json['cityPostalCodes'].map(CityPostalCodeType_1.CityPostalCodeTypeFromJSON)),
        'masterInfoList': !(0, runtime_1.exists)(json, 'masterInfoList') ? undefined : (json['masterInfoList'].map(MasterInfoType_1.MasterInfoTypeFromJSON)),
        'totalPages': !(0, runtime_1.exists)(json, 'totalPages') ? undefined : json['totalPages'],
        'offset': !(0, runtime_1.exists)(json, 'offset') ? undefined : json['offset'],
        'limit': !(0, runtime_1.exists)(json, 'limit') ? undefined : json['limit'],
        'hasMore': !(0, runtime_1.exists)(json, 'hasMore') ? undefined : json['hasMore'],
        'totalResults': !(0, runtime_1.exists)(json, 'totalResults') ? undefined : json['totalResults'],
        'count': !(0, runtime_1.exists)(json, 'count') ? undefined : json['count'],
        'links': !(0, runtime_1.exists)(json, 'links') ? undefined : (json['links'].map(InstanceLink_1.InstanceLinkFromJSON)),
        'warnings': !(0, runtime_1.exists)(json, 'warnings') ? undefined : (json['warnings'].map(WarningType_1.WarningTypeFromJSON)),
    };
}
exports.CityPostalCodesDetailsFromJSONTyped = CityPostalCodesDetailsFromJSONTyped;
function CityPostalCodesDetailsToJSON(value) {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        'cityPostalCodes': value.cityPostalCodes === undefined ? undefined : (value.cityPostalCodes.map(CityPostalCodeType_1.CityPostalCodeTypeToJSON)),
        'masterInfoList': value.masterInfoList === undefined ? undefined : (value.masterInfoList.map(MasterInfoType_1.MasterInfoTypeToJSON)),
        'totalPages': value.totalPages,
        'offset': value.offset,
        'limit': value.limit,
        'hasMore': value.hasMore,
        'totalResults': value.totalResults,
        'count': value.count,
        'links': value.links === undefined ? undefined : (value.links.map(InstanceLink_1.InstanceLinkToJSON)),
        'warnings': value.warnings === undefined ? undefined : (value.warnings.map(WarningType_1.WarningTypeToJSON)),
    };
}
exports.CityPostalCodesDetailsToJSON = CityPostalCodesDetailsToJSON;
