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
exports.CopyActivityLocationsRSToJSON = exports.CopyActivityLocationsRSFromJSONTyped = exports.CopyActivityLocationsRSFromJSON = exports.instanceOfCopyActivityLocationsRS = void 0;
const runtime_1 = require("../runtime");
const ErrorType_1 = require("./ErrorType");
const InstanceLink_1 = require("./InstanceLink");
const WarningType_1 = require("./WarningType");
/**
 * Check if a given object implements the CopyActivityLocationsRS interface.
 */
function instanceOfCopyActivityLocationsRS(value) {
    let isInstance = true;
    return isInstance;
}
exports.instanceOfCopyActivityLocationsRS = instanceOfCopyActivityLocationsRS;
function CopyActivityLocationsRSFromJSON(json) {
    return CopyActivityLocationsRSFromJSONTyped(json, false);
}
exports.CopyActivityLocationsRSFromJSON = CopyActivityLocationsRSFromJSON;
function CopyActivityLocationsRSFromJSONTyped(json, ignoreDiscriminator) {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        'success': !(0, runtime_1.exists)(json, 'success') ? undefined : json['success'],
        'warnings': !(0, runtime_1.exists)(json, 'warnings') ? undefined : (json['warnings'].map(WarningType_1.WarningTypeFromJSON)),
        'errors': !(0, runtime_1.exists)(json, 'errors') ? undefined : (json['errors'].map(ErrorType_1.ErrorTypeFromJSON)),
        'echoToken': !(0, runtime_1.exists)(json, 'echoToken') ? undefined : json['echoToken'],
        'timeStamp': !(0, runtime_1.exists)(json, 'timeStamp') ? undefined : (new Date(json['timeStamp'])),
        'version': !(0, runtime_1.exists)(json, 'version') ? undefined : json['version'],
        'correlationId': !(0, runtime_1.exists)(json, 'correlationId') ? undefined : json['correlationId'],
        'retryAllowed': !(0, runtime_1.exists)(json, 'retryAllowed') ? undefined : json['retryAllowed'],
        'enforceAllowed': !(0, runtime_1.exists)(json, 'enforceAllowed') ? undefined : json['enforceAllowed'],
        'useLocalAllowed': !(0, runtime_1.exists)(json, 'useLocalAllowed') ? undefined : json['useLocalAllowed'],
        'links': !(0, runtime_1.exists)(json, 'links') ? undefined : (json['links'].map(InstanceLink_1.InstanceLinkFromJSON)),
    };
}
exports.CopyActivityLocationsRSFromJSONTyped = CopyActivityLocationsRSFromJSONTyped;
function CopyActivityLocationsRSToJSON(value) {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        'success': value.success,
        'warnings': value.warnings === undefined ? undefined : (value.warnings.map(WarningType_1.WarningTypeToJSON)),
        'errors': value.errors === undefined ? undefined : (value.errors.map(ErrorType_1.ErrorTypeToJSON)),
        'echoToken': value.echoToken,
        'timeStamp': value.timeStamp === undefined ? undefined : (value.timeStamp.toISOString()),
        'version': value.version,
        'correlationId': value.correlationId,
        'retryAllowed': value.retryAllowed,
        'enforceAllowed': value.enforceAllowed,
        'useLocalAllowed': value.useLocalAllowed,
        'links': value.links === undefined ? undefined : (value.links.map(InstanceLink_1.InstanceLinkToJSON)),
    };
}
exports.CopyActivityLocationsRSToJSON = CopyActivityLocationsRSToJSON;
