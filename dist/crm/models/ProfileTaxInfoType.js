"use strict";
/* tslint:disable */
/* eslint-disable */
/**
 * OPERA Cloud Customer Relationship Management API
 * APIs to cater for Customer Relationship Management (profile) functionality in OPERA Cloud.  There are different types of profiles in OPERA Cloud, including Guest, Company, Travel Agent, Source, Group, and Contact profile types.  A profile can store and display a wide range of information about the guest, company, travel agent etc., depending upon the type of profile.  For example, a guest profile can store the guest name, address, contact information, details on billing, membership benefits, preferences and much more.  All profiles in OPERA when created are assigned a ProfileID.  This ID will be used throughout the CRM APIs.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 21.5.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
Object.defineProperty(exports, "__esModule", { value: true });
exports.ProfileTaxInfoTypeToJSON = exports.ProfileTaxInfoTypeFromJSONTyped = exports.ProfileTaxInfoTypeFromJSON = exports.instanceOfProfileTaxInfoType = void 0;
const runtime_1 = require("../runtime");
/**
 * Check if a given object implements the ProfileTaxInfoType interface.
 */
function instanceOfProfileTaxInfoType(value) {
    let isInstance = true;
    return isInstance;
}
exports.instanceOfProfileTaxInfoType = instanceOfProfileTaxInfoType;
function ProfileTaxInfoTypeFromJSON(json) {
    return ProfileTaxInfoTypeFromJSONTyped(json, false);
}
exports.ProfileTaxInfoTypeFromJSON = ProfileTaxInfoTypeFromJSON;
function ProfileTaxInfoTypeFromJSONTyped(json, ignoreDiscriminator) {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        'tax1No': !(0, runtime_1.exists)(json, 'tax1No') ? undefined : json['tax1No'],
        'tax2No': !(0, runtime_1.exists)(json, 'tax2No') ? undefined : json['tax2No'],
        'taxCategory': !(0, runtime_1.exists)(json, 'taxCategory') ? undefined : json['taxCategory'],
        'taxOffice': !(0, runtime_1.exists)(json, 'taxOffice') ? undefined : json['taxOffice'],
        'taxType': !(0, runtime_1.exists)(json, 'taxType') ? undefined : json['taxType'],
        'businessId': !(0, runtime_1.exists)(json, 'businessId') ? undefined : json['businessId'],
        'businessRegistration': !(0, runtime_1.exists)(json, 'businessRegistration') ? undefined : json['businessRegistration'],
    };
}
exports.ProfileTaxInfoTypeFromJSONTyped = ProfileTaxInfoTypeFromJSONTyped;
function ProfileTaxInfoTypeToJSON(value) {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        'tax1No': value.tax1No,
        'tax2No': value.tax2No,
        'taxCategory': value.taxCategory,
        'taxOffice': value.taxOffice,
        'taxType': value.taxType,
        'businessId': value.businessId,
        'businessRegistration': value.businessRegistration,
    };
}
exports.ProfileTaxInfoTypeToJSON = ProfileTaxInfoTypeToJSON;
