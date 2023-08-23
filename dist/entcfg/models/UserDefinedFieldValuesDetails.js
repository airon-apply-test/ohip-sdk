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
exports.UserDefinedFieldValuesDetailsToJSON = exports.UserDefinedFieldValuesDetailsFromJSONTyped = exports.UserDefinedFieldValuesDetailsFromJSON = exports.instanceOfUserDefinedFieldValuesDetails = void 0;
const runtime_1 = require("../runtime");
const InstanceLink_1 = require("./InstanceLink");
const UserDefinedFieldValueType_1 = require("./UserDefinedFieldValueType");
const WarningType_1 = require("./WarningType");
/**
 * Check if a given object implements the UserDefinedFieldValuesDetails interface.
 */
function instanceOfUserDefinedFieldValuesDetails(value) {
    let isInstance = true;
    return isInstance;
}
exports.instanceOfUserDefinedFieldValuesDetails = instanceOfUserDefinedFieldValuesDetails;
function UserDefinedFieldValuesDetailsFromJSON(json) {
    return UserDefinedFieldValuesDetailsFromJSONTyped(json, false);
}
exports.UserDefinedFieldValuesDetailsFromJSON = UserDefinedFieldValuesDetailsFromJSON;
function UserDefinedFieldValuesDetailsFromJSONTyped(json, ignoreDiscriminator) {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        'userDefinedFieldValues': !(0, runtime_1.exists)(json, 'userDefinedFieldValues') ? undefined : (json['userDefinedFieldValues'].map(UserDefinedFieldValueType_1.UserDefinedFieldValueTypeFromJSON)),
        'links': !(0, runtime_1.exists)(json, 'links') ? undefined : (json['links'].map(InstanceLink_1.InstanceLinkFromJSON)),
        'warnings': !(0, runtime_1.exists)(json, 'warnings') ? undefined : (json['warnings'].map(WarningType_1.WarningTypeFromJSON)),
    };
}
exports.UserDefinedFieldValuesDetailsFromJSONTyped = UserDefinedFieldValuesDetailsFromJSONTyped;
function UserDefinedFieldValuesDetailsToJSON(value) {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        'userDefinedFieldValues': value.userDefinedFieldValues === undefined ? undefined : (value.userDefinedFieldValues.map(UserDefinedFieldValueType_1.UserDefinedFieldValueTypeToJSON)),
        'links': value.links === undefined ? undefined : (value.links.map(InstanceLink_1.InstanceLinkToJSON)),
        'warnings': value.warnings === undefined ? undefined : (value.warnings.map(WarningType_1.WarningTypeToJSON)),
    };
}
exports.UserDefinedFieldValuesDetailsToJSON = UserDefinedFieldValuesDetailsToJSON;
