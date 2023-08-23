"use strict";
/* tslint:disable */
/* eslint-disable */
/**
 * OPERA Cloud Integration Configuration API
 * APIs catering to Integration Configuration in OPERA Cloud.  Operations such as get Hotel Interface Types, or get UDF mappings can be found in this module.<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 22.3.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
Object.defineProperty(exports, "__esModule", { value: true });
exports.FetchedUDFMappingsToJSON = exports.FetchedUDFMappingsFromJSONTyped = exports.FetchedUDFMappingsFromJSON = exports.instanceOfFetchedUDFMappings = void 0;
const runtime_1 = require("../runtime");
const InstanceLink_1 = require("./InstanceLink");
const UDFMappingType_1 = require("./UDFMappingType");
const WarningType_1 = require("./WarningType");
/**
 * Check if a given object implements the FetchedUDFMappings interface.
 */
function instanceOfFetchedUDFMappings(value) {
    let isInstance = true;
    return isInstance;
}
exports.instanceOfFetchedUDFMappings = instanceOfFetchedUDFMappings;
function FetchedUDFMappingsFromJSON(json) {
    return FetchedUDFMappingsFromJSONTyped(json, false);
}
exports.FetchedUDFMappingsFromJSON = FetchedUDFMappingsFromJSON;
function FetchedUDFMappingsFromJSONTyped(json, ignoreDiscriminator) {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        'uDFMappings': !(0, runtime_1.exists)(json, 'uDFMappings') ? undefined : (json['uDFMappings'].map(UDFMappingType_1.UDFMappingTypeFromJSON)),
        'links': !(0, runtime_1.exists)(json, 'links') ? undefined : (json['links'].map(InstanceLink_1.InstanceLinkFromJSON)),
        'warnings': !(0, runtime_1.exists)(json, 'warnings') ? undefined : (json['warnings'].map(WarningType_1.WarningTypeFromJSON)),
    };
}
exports.FetchedUDFMappingsFromJSONTyped = FetchedUDFMappingsFromJSONTyped;
function FetchedUDFMappingsToJSON(value) {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        'uDFMappings': value.uDFMappings === undefined ? undefined : (value.uDFMappings.map(UDFMappingType_1.UDFMappingTypeToJSON)),
        'links': value.links === undefined ? undefined : (value.links.map(InstanceLink_1.InstanceLinkToJSON)),
        'warnings': value.warnings === undefined ? undefined : (value.warnings.map(WarningType_1.WarningTypeToJSON)),
    };
}
exports.FetchedUDFMappingsToJSON = FetchedUDFMappingsToJSON;
