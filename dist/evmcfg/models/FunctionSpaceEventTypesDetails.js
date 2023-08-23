"use strict";
/* tslint:disable */
/* eslint-disable */
/**
 * OPERA Cloud Event Configuration API
 * This API caters for Event Configuration in OPERA Cloud. <br /><There are operations to post, update, fetch and delete codes such as item inventory, function spaces, menu items and many more.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 21.5.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
Object.defineProperty(exports, "__esModule", { value: true });
exports.FunctionSpaceEventTypesDetailsToJSON = exports.FunctionSpaceEventTypesDetailsFromJSONTyped = exports.FunctionSpaceEventTypesDetailsFromJSON = exports.instanceOfFunctionSpaceEventTypesDetails = void 0;
const runtime_1 = require("../runtime");
const EventTypeConfigType_1 = require("./EventTypeConfigType");
const InstanceLink_1 = require("./InstanceLink");
const WarningType_1 = require("./WarningType");
/**
 * Check if a given object implements the FunctionSpaceEventTypesDetails interface.
 */
function instanceOfFunctionSpaceEventTypesDetails(value) {
    let isInstance = true;
    return isInstance;
}
exports.instanceOfFunctionSpaceEventTypesDetails = instanceOfFunctionSpaceEventTypesDetails;
function FunctionSpaceEventTypesDetailsFromJSON(json) {
    return FunctionSpaceEventTypesDetailsFromJSONTyped(json, false);
}
exports.FunctionSpaceEventTypesDetailsFromJSON = FunctionSpaceEventTypesDetailsFromJSON;
function FunctionSpaceEventTypesDetailsFromJSONTyped(json, ignoreDiscriminator) {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        'eventTypes': !(0, runtime_1.exists)(json, 'eventTypes') ? undefined : (json['eventTypes'].map(EventTypeConfigType_1.EventTypeConfigTypeFromJSON)),
        'links': !(0, runtime_1.exists)(json, 'links') ? undefined : (json['links'].map(InstanceLink_1.InstanceLinkFromJSON)),
        'warnings': !(0, runtime_1.exists)(json, 'warnings') ? undefined : (json['warnings'].map(WarningType_1.WarningTypeFromJSON)),
    };
}
exports.FunctionSpaceEventTypesDetailsFromJSONTyped = FunctionSpaceEventTypesDetailsFromJSONTyped;
function FunctionSpaceEventTypesDetailsToJSON(value) {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        'eventTypes': value.eventTypes === undefined ? undefined : (value.eventTypes.map(EventTypeConfigType_1.EventTypeConfigTypeToJSON)),
        'links': value.links === undefined ? undefined : (value.links.map(InstanceLink_1.InstanceLinkToJSON)),
        'warnings': value.warnings === undefined ? undefined : (value.warnings.map(WarningType_1.WarningTypeToJSON)),
    };
}
exports.FunctionSpaceEventTypesDetailsToJSON = FunctionSpaceEventTypesDetailsToJSON;
