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
exports.ProformaResponseInstructionTypeToJSON = exports.ProformaResponseInstructionTypeFromJSONTyped = exports.ProformaResponseInstructionTypeFromJSON = exports.ProformaResponseInstructionType = void 0;
/**
 * Response instructions for Proforma Creation.
 * @export
 */
exports.ProformaResponseInstructionType = {
    None: 'None',
    FolioSummary: 'FolioSummary'
};
function ProformaResponseInstructionTypeFromJSON(json) {
    return ProformaResponseInstructionTypeFromJSONTyped(json, false);
}
exports.ProformaResponseInstructionTypeFromJSON = ProformaResponseInstructionTypeFromJSON;
function ProformaResponseInstructionTypeFromJSONTyped(json, ignoreDiscriminator) {
    return json;
}
exports.ProformaResponseInstructionTypeFromJSONTyped = ProformaResponseInstructionTypeFromJSONTyped;
function ProformaResponseInstructionTypeToJSON(value) {
    return value;
}
exports.ProformaResponseInstructionTypeToJSON = ProformaResponseInstructionTypeToJSON;
