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
exports.FiscalResponseTypeToJSON = exports.FiscalResponseTypeFromJSONTyped = exports.FiscalResponseTypeFromJSON = exports.instanceOfFiscalResponseType = void 0;
const runtime_1 = require("../runtime");
/**
 * Check if a given object implements the FiscalResponseType interface.
 */
function instanceOfFiscalResponseType(value) {
    let isInstance = true;
    return isInstance;
}
exports.instanceOfFiscalResponseType = instanceOfFiscalResponseType;
function FiscalResponseTypeFromJSON(json) {
    return FiscalResponseTypeFromJSONTyped(json, false);
}
exports.FiscalResponseTypeFromJSON = FiscalResponseTypeFromJSON;
function FiscalResponseTypeFromJSONTyped(json, ignoreDiscriminator) {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        'folioSeqId': !(0, runtime_1.exists)(json, 'folioSeqId') ? undefined : json['folioSeqId'],
        'retryFiscalPrinting': !(0, runtime_1.exists)(json, 'retryFiscalPrinting') ? undefined : json['retryFiscalPrinting'],
    };
}
exports.FiscalResponseTypeFromJSONTyped = FiscalResponseTypeFromJSONTyped;
function FiscalResponseTypeToJSON(value) {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        'folioSeqId': value.folioSeqId,
        'retryFiscalPrinting': value.retryFiscalPrinting,
    };
}
exports.FiscalResponseTypeToJSON = FiscalResponseTypeToJSON;
