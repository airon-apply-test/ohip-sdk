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
exports.ServingTypeToJSON = exports.ServingTypeFromJSONTyped = exports.ServingTypeFromJSON = exports.ServingType = void 0;
/**
 * Serving
 * @export
 */
exports.ServingType = {
    Pp: 'Pp',
    Pt: 'Pt'
};
function ServingTypeFromJSON(json) {
    return ServingTypeFromJSONTyped(json, false);
}
exports.ServingTypeFromJSON = ServingTypeFromJSON;
function ServingTypeFromJSONTyped(json, ignoreDiscriminator) {
    return json;
}
exports.ServingTypeFromJSONTyped = ServingTypeFromJSONTyped;
function ServingTypeToJSON(value) {
    return value;
}
exports.ServingTypeToJSON = ServingTypeToJSON;
