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
exports.PutActivityStatusCodesRequestToJSON = exports.PutActivityStatusCodesRequestFromJSONTyped = exports.PutActivityStatusCodesRequestFromJSON = exports.instanceOfPutActivityStatusCodesRequest = void 0;
const runtime_1 = require("../runtime");
const ActivityStatusCodeType_1 = require("./ActivityStatusCodeType");
const InstanceLink_1 = require("./InstanceLink");
const WarningType_1 = require("./WarningType");
/**
 * Check if a given object implements the PutActivityStatusCodesRequest interface.
 */
function instanceOfPutActivityStatusCodesRequest(value) {
    let isInstance = true;
    return isInstance;
}
exports.instanceOfPutActivityStatusCodesRequest = instanceOfPutActivityStatusCodesRequest;
function PutActivityStatusCodesRequestFromJSON(json) {
    return PutActivityStatusCodesRequestFromJSONTyped(json, false);
}
exports.PutActivityStatusCodesRequestFromJSON = PutActivityStatusCodesRequestFromJSON;
function PutActivityStatusCodesRequestFromJSONTyped(json, ignoreDiscriminator) {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        'activityStatusCodes': !(0, runtime_1.exists)(json, 'activityStatusCodes') ? undefined : (json['activityStatusCodes'].map(ActivityStatusCodeType_1.ActivityStatusCodeTypeFromJSON)),
        'links': !(0, runtime_1.exists)(json, 'links') ? undefined : (json['links'].map(InstanceLink_1.InstanceLinkFromJSON)),
        'warnings': !(0, runtime_1.exists)(json, 'warnings') ? undefined : (json['warnings'].map(WarningType_1.WarningTypeFromJSON)),
    };
}
exports.PutActivityStatusCodesRequestFromJSONTyped = PutActivityStatusCodesRequestFromJSONTyped;
function PutActivityStatusCodesRequestToJSON(value) {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        'activityStatusCodes': value.activityStatusCodes === undefined ? undefined : (value.activityStatusCodes.map(ActivityStatusCodeType_1.ActivityStatusCodeTypeToJSON)),
        'links': value.links === undefined ? undefined : (value.links.map(InstanceLink_1.InstanceLinkToJSON)),
        'warnings': value.warnings === undefined ? undefined : (value.warnings.map(WarningType_1.WarningTypeToJSON)),
    };
}
exports.PutActivityStatusCodesRequestToJSON = PutActivityStatusCodesRequestToJSON;
