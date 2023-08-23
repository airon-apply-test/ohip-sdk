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
exports.FetchedHotelInterfaceSchemasToJSON = exports.FetchedHotelInterfaceSchemasFromJSONTyped = exports.FetchedHotelInterfaceSchemasFromJSON = exports.instanceOfFetchedHotelInterfaceSchemas = void 0;
const runtime_1 = require("../runtime");
const HotelInterfaceSchemaType_1 = require("./HotelInterfaceSchemaType");
const InstanceLink_1 = require("./InstanceLink");
const WarningType_1 = require("./WarningType");
/**
 * Check if a given object implements the FetchedHotelInterfaceSchemas interface.
 */
function instanceOfFetchedHotelInterfaceSchemas(value) {
    let isInstance = true;
    return isInstance;
}
exports.instanceOfFetchedHotelInterfaceSchemas = instanceOfFetchedHotelInterfaceSchemas;
function FetchedHotelInterfaceSchemasFromJSON(json) {
    return FetchedHotelInterfaceSchemasFromJSONTyped(json, false);
}
exports.FetchedHotelInterfaceSchemasFromJSON = FetchedHotelInterfaceSchemasFromJSON;
function FetchedHotelInterfaceSchemasFromJSONTyped(json, ignoreDiscriminator) {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        'schemas': !(0, runtime_1.exists)(json, 'schemas') ? undefined : (json['schemas'].map(HotelInterfaceSchemaType_1.HotelInterfaceSchemaTypeFromJSON)),
        'totalPages': !(0, runtime_1.exists)(json, 'totalPages') ? undefined : json['totalPages'],
        'offset': !(0, runtime_1.exists)(json, 'offset') ? undefined : json['offset'],
        'limit': !(0, runtime_1.exists)(json, 'limit') ? undefined : json['limit'],
        'hasMore': !(0, runtime_1.exists)(json, 'hasMore') ? undefined : json['hasMore'],
        'totalResults': !(0, runtime_1.exists)(json, 'totalResults') ? undefined : json['totalResults'],
        'links': !(0, runtime_1.exists)(json, 'links') ? undefined : (json['links'].map(InstanceLink_1.InstanceLinkFromJSON)),
        'warnings': !(0, runtime_1.exists)(json, 'warnings') ? undefined : (json['warnings'].map(WarningType_1.WarningTypeFromJSON)),
    };
}
exports.FetchedHotelInterfaceSchemasFromJSONTyped = FetchedHotelInterfaceSchemasFromJSONTyped;
function FetchedHotelInterfaceSchemasToJSON(value) {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        'schemas': value.schemas === undefined ? undefined : (value.schemas.map(HotelInterfaceSchemaType_1.HotelInterfaceSchemaTypeToJSON)),
        'totalPages': value.totalPages,
        'offset': value.offset,
        'limit': value.limit,
        'hasMore': value.hasMore,
        'totalResults': value.totalResults,
        'links': value.links === undefined ? undefined : (value.links.map(InstanceLink_1.InstanceLinkToJSON)),
        'warnings': value.warnings === undefined ? undefined : (value.warnings.map(WarningType_1.WarningTypeToJSON)),
    };
}
exports.FetchedHotelInterfaceSchemasToJSON = FetchedHotelInterfaceSchemasToJSON;
