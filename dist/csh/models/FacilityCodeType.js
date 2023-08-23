"use strict";
/* tslint:disable */
/* eslint-disable */
/**
 * OPERA Cloud Cashiering API
 * APIs to cater for Cashiering related functionality in OPERA Cloud.<br /><br /> Cashiering provides access to a guest folio, posting journals, receipt histories, currency calculations, credit card settlements, and check a guest out.<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 22.3.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
Object.defineProperty(exports, "__esModule", { value: true });
exports.FacilityCodeTypeToJSON = exports.FacilityCodeTypeFromJSONTyped = exports.FacilityCodeTypeFromJSON = exports.instanceOfFacilityCodeType = void 0;
const runtime_1 = require("../runtime");
/**
 * Check if a given object implements the FacilityCodeType interface.
 */
function instanceOfFacilityCodeType(value) {
    let isInstance = true;
    return isInstance;
}
exports.instanceOfFacilityCodeType = instanceOfFacilityCodeType;
function FacilityCodeTypeFromJSON(json) {
    return FacilityCodeTypeFromJSONTyped(json, false);
}
exports.FacilityCodeTypeFromJSON = FacilityCodeTypeFromJSON;
function FacilityCodeTypeFromJSONTyped(json, ignoreDiscriminator) {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        'description': !(0, runtime_1.exists)(json, 'description') ? undefined : json['description'],
        'quantity': !(0, runtime_1.exists)(json, 'quantity') ? undefined : json['quantity'],
        'code': !(0, runtime_1.exists)(json, 'code') ? undefined : json['code'],
    };
}
exports.FacilityCodeTypeFromJSONTyped = FacilityCodeTypeFromJSONTyped;
function FacilityCodeTypeToJSON(value) {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        'description': value.description,
        'quantity': value.quantity,
        'code': value.code,
    };
}
exports.FacilityCodeTypeToJSON = FacilityCodeTypeToJSON;
