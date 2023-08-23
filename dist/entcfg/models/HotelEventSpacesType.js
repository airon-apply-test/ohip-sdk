"use strict";
/* tslint:disable */
/* eslint-disable */
/**
 * OPERA Cloud Enterprise Configuration API
 * APIs to cater for Enterprise Configuration functionality in OPERA Cloud. <br /><br In this module, you can configure a variety of options related to your properties such as their locations, facilities, and local attractions. The available options are dependant on the active controls at your property.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 21.5.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
Object.defineProperty(exports, "__esModule", { value: true });
exports.HotelEventSpacesTypeToJSON = exports.HotelEventSpacesTypeFromJSONTyped = exports.HotelEventSpacesTypeFromJSON = exports.instanceOfHotelEventSpacesType = void 0;
const runtime_1 = require("../runtime");
const HotelEventSpaceDetailType_1 = require("./HotelEventSpaceDetailType");
const HotelEventSpaceSummaryType_1 = require("./HotelEventSpaceSummaryType");
/**
 * Check if a given object implements the HotelEventSpacesType interface.
 */
function instanceOfHotelEventSpacesType(value) {
    let isInstance = true;
    return isInstance;
}
exports.instanceOfHotelEventSpacesType = instanceOfHotelEventSpacesType;
function HotelEventSpacesTypeFromJSON(json) {
    return HotelEventSpacesTypeFromJSONTyped(json, false);
}
exports.HotelEventSpacesTypeFromJSON = HotelEventSpacesTypeFromJSON;
function HotelEventSpacesTypeFromJSONTyped(json, ignoreDiscriminator) {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        'eventSpaceDetails': !(0, runtime_1.exists)(json, 'eventSpaceDetails') ? undefined : (json['eventSpaceDetails'].map(HotelEventSpaceDetailType_1.HotelEventSpaceDetailTypeFromJSON)),
        'eventSpaceSummaries': !(0, runtime_1.exists)(json, 'eventSpaceSummaries') ? undefined : (json['eventSpaceSummaries'].map(HotelEventSpaceSummaryType_1.HotelEventSpaceSummaryTypeFromJSON)),
        'setupStyles': !(0, runtime_1.exists)(json, 'setupStyles') ? undefined : json['setupStyles'],
    };
}
exports.HotelEventSpacesTypeFromJSONTyped = HotelEventSpacesTypeFromJSONTyped;
function HotelEventSpacesTypeToJSON(value) {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        'eventSpaceDetails': value.eventSpaceDetails === undefined ? undefined : (value.eventSpaceDetails.map(HotelEventSpaceDetailType_1.HotelEventSpaceDetailTypeToJSON)),
        'eventSpaceSummaries': value.eventSpaceSummaries === undefined ? undefined : (value.eventSpaceSummaries.map(HotelEventSpaceSummaryType_1.HotelEventSpaceSummaryTypeToJSON)),
        'setupStyles': value.setupStyles,
    };
}
exports.HotelEventSpacesTypeToJSON = HotelEventSpacesTypeToJSON;
