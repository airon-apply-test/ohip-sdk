"use strict";
/* tslint:disable */
/* eslint-disable */
/**
 * OPERA Cloud Leisure Management API
 * APIs to cater for external Leisure Management functionality integrated with OPERA Cloud.<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 22.3.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
Object.defineProperty(exports, "__esModule", { value: true });
exports.FetchActivityTypesToJSON = exports.FetchActivityTypesFromJSONTyped = exports.FetchActivityTypesFromJSON = exports.instanceOfFetchActivityTypes = void 0;
const runtime_1 = require("../runtime");
const ActivityTypeDetailType_1 = require("./ActivityTypeDetailType");
const InstanceLink_1 = require("./InstanceLink");
const WarningType_1 = require("./WarningType");
/**
 * Check if a given object implements the FetchActivityTypes interface.
 */
function instanceOfFetchActivityTypes(value) {
    let isInstance = true;
    return isInstance;
}
exports.instanceOfFetchActivityTypes = instanceOfFetchActivityTypes;
function FetchActivityTypesFromJSON(json) {
    return FetchActivityTypesFromJSONTyped(json, false);
}
exports.FetchActivityTypesFromJSON = FetchActivityTypesFromJSON;
function FetchActivityTypesFromJSONTyped(json, ignoreDiscriminator) {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        'activityTypes': !(0, runtime_1.exists)(json, 'activityTypes') ? undefined : (json['activityTypes'].map(ActivityTypeDetailType_1.ActivityTypeDetailTypeFromJSON)),
        'links': !(0, runtime_1.exists)(json, 'links') ? undefined : (json['links'].map(InstanceLink_1.InstanceLinkFromJSON)),
        'warnings': !(0, runtime_1.exists)(json, 'warnings') ? undefined : (json['warnings'].map(WarningType_1.WarningTypeFromJSON)),
    };
}
exports.FetchActivityTypesFromJSONTyped = FetchActivityTypesFromJSONTyped;
function FetchActivityTypesToJSON(value) {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        'activityTypes': value.activityTypes === undefined ? undefined : (value.activityTypes.map(ActivityTypeDetailType_1.ActivityTypeDetailTypeToJSON)),
        'links': value.links === undefined ? undefined : (value.links.map(InstanceLink_1.InstanceLinkToJSON)),
        'warnings': value.warnings === undefined ? undefined : (value.warnings.map(WarningType_1.WarningTypeToJSON)),
    };
}
exports.FetchActivityTypesToJSON = FetchActivityTypesToJSON;
