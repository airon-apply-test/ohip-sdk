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
exports.RedeemPrepaidCardsStatusToJSON = exports.RedeemPrepaidCardsStatusFromJSONTyped = exports.RedeemPrepaidCardsStatusFromJSON = exports.instanceOfRedeemPrepaidCardsStatus = void 0;
const runtime_1 = require("../runtime");
const PrepaidCardInfoType_1 = require("./PrepaidCardInfoType");
const WarningType_1 = require("./WarningType");
/**
 * Check if a given object implements the RedeemPrepaidCardsStatus interface.
 */
function instanceOfRedeemPrepaidCardsStatus(value) {
    let isInstance = true;
    return isInstance;
}
exports.instanceOfRedeemPrepaidCardsStatus = instanceOfRedeemPrepaidCardsStatus;
function RedeemPrepaidCardsStatusFromJSON(json) {
    return RedeemPrepaidCardsStatusFromJSONTyped(json, false);
}
exports.RedeemPrepaidCardsStatusFromJSON = RedeemPrepaidCardsStatusFromJSON;
function RedeemPrepaidCardsStatusFromJSONTyped(json, ignoreDiscriminator) {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        'prepaidCards': !(0, runtime_1.exists)(json, 'prepaidCards') ? undefined : (json['prepaidCards'].map(PrepaidCardInfoType_1.PrepaidCardInfoTypeFromJSON)),
        'warnings': !(0, runtime_1.exists)(json, 'warnings') ? undefined : (json['warnings'].map(WarningType_1.WarningTypeFromJSON)),
    };
}
exports.RedeemPrepaidCardsStatusFromJSONTyped = RedeemPrepaidCardsStatusFromJSONTyped;
function RedeemPrepaidCardsStatusToJSON(value) {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        'prepaidCards': value.prepaidCards === undefined ? undefined : (value.prepaidCards.map(PrepaidCardInfoType_1.PrepaidCardInfoTypeToJSON)),
        'warnings': value.warnings === undefined ? undefined : (value.warnings.map(WarningType_1.WarningTypeToJSON)),
    };
}
exports.RedeemPrepaidCardsStatusToJSON = RedeemPrepaidCardsStatusToJSON;
