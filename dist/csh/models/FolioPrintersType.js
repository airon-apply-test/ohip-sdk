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
exports.FolioPrintersTypeToJSON = exports.FolioPrintersTypeFromJSONTyped = exports.FolioPrintersTypeFromJSON = exports.instanceOfFolioPrintersType = void 0;
const runtime_1 = require("../runtime");
const FolioPrinterType_1 = require("./FolioPrinterType");
/**
 * Check if a given object implements the FolioPrintersType interface.
 */
function instanceOfFolioPrintersType(value) {
    let isInstance = true;
    return isInstance;
}
exports.instanceOfFolioPrintersType = instanceOfFolioPrintersType;
function FolioPrintersTypeFromJSON(json) {
    return FolioPrintersTypeFromJSONTyped(json, false);
}
exports.FolioPrintersTypeFromJSON = FolioPrintersTypeFromJSON;
function FolioPrintersTypeFromJSONTyped(json, ignoreDiscriminator) {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        'folioTypeName': !(0, runtime_1.exists)(json, 'folioTypeName') ? undefined : json['folioTypeName'],
        'printer': !(0, runtime_1.exists)(json, 'printer') ? undefined : (json['printer'].map(FolioPrinterType_1.FolioPrinterTypeFromJSON)),
    };
}
exports.FolioPrintersTypeFromJSONTyped = FolioPrintersTypeFromJSONTyped;
function FolioPrintersTypeToJSON(value) {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        'folioTypeName': value.folioTypeName,
        'printer': value.printer === undefined ? undefined : (value.printer.map(FolioPrinterType_1.FolioPrinterTypeToJSON)),
    };
}
exports.FolioPrintersTypeToJSON = FolioPrintersTypeToJSON;
