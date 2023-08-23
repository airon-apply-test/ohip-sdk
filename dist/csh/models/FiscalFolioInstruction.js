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
exports.FiscalFolioInstructionToJSON = exports.FiscalFolioInstructionFromJSONTyped = exports.FiscalFolioInstructionFromJSON = exports.FiscalFolioInstruction = void 0;
/**
 * Action to generate Offline folio when no reponse is received from fiscal service.
 * @export
 */
exports.FiscalFolioInstruction = {
    Retry: 'Retry',
    New: 'New',
    Void: 'Void',
    Offline: 'Offline'
};
function FiscalFolioInstructionFromJSON(json) {
    return FiscalFolioInstructionFromJSONTyped(json, false);
}
exports.FiscalFolioInstructionFromJSON = FiscalFolioInstructionFromJSON;
function FiscalFolioInstructionFromJSONTyped(json, ignoreDiscriminator) {
    return json;
}
exports.FiscalFolioInstructionFromJSONTyped = FiscalFolioInstructionFromJSONTyped;
function FiscalFolioInstructionToJSON(value) {
    return value;
}
exports.FiscalFolioInstructionToJSON = FiscalFolioInstructionToJSON;
