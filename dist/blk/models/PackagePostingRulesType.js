"use strict";
/* tslint:disable */
/* eslint-disable */
/**
 * OPERA Cloud Block API
 * APIs to cater for Business Block functionality in OPERA Cloud. <br /><br /> A block is a group of rooms held for guests who are attending an event, meeting, or function. You can create blocks for family reunions, business conferences, weddings, and so on. You can also set aside rooms for the event (block).<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 22.3.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
Object.defineProperty(exports, "__esModule", { value: true });
exports.PackagePostingRulesTypeToJSON = exports.PackagePostingRulesTypeFromJSONTyped = exports.PackagePostingRulesTypeFromJSON = exports.instanceOfPackagePostingRulesType = void 0;
const runtime_1 = require("../runtime");
const PackageTransactionCodeType_1 = require("./PackageTransactionCodeType");
/**
 * Check if a given object implements the PackagePostingRulesType interface.
 */
function instanceOfPackagePostingRulesType(value) {
    let isInstance = true;
    return isInstance;
}
exports.instanceOfPackagePostingRulesType = instanceOfPackagePostingRulesType;
function PackagePostingRulesTypeFromJSON(json) {
    return PackagePostingRulesTypeFromJSONTyped(json, false);
}
exports.PackagePostingRulesTypeFromJSON = PackagePostingRulesTypeFromJSON;
function PackagePostingRulesTypeFromJSONTyped(json, ignoreDiscriminator) {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        'transactionCode': !(0, runtime_1.exists)(json, 'transactionCode') ? undefined : (0, PackageTransactionCodeType_1.PackageTransactionCodeTypeFromJSON)(json['transactionCode']),
        'overageCode': !(0, runtime_1.exists)(json, 'overageCode') ? undefined : (0, PackageTransactionCodeType_1.PackageTransactionCodeTypeFromJSON)(json['overageCode']),
        'profitCode': !(0, runtime_1.exists)(json, 'profitCode') ? undefined : (0, PackageTransactionCodeType_1.PackageTransactionCodeTypeFromJSON)(json['profitCode']),
        'lossCode': !(0, runtime_1.exists)(json, 'lossCode') ? undefined : (0, PackageTransactionCodeType_1.PackageTransactionCodeTypeFromJSON)(json['lossCode']),
        'alternateCodesList': !(0, runtime_1.exists)(json, 'alternateCodesList') ? undefined : (json['alternateCodesList'].map(PackageTransactionCodeType_1.PackageTransactionCodeTypeFromJSON)),
    };
}
exports.PackagePostingRulesTypeFromJSONTyped = PackagePostingRulesTypeFromJSONTyped;
function PackagePostingRulesTypeToJSON(value) {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        'transactionCode': (0, PackageTransactionCodeType_1.PackageTransactionCodeTypeToJSON)(value.transactionCode),
        'overageCode': (0, PackageTransactionCodeType_1.PackageTransactionCodeTypeToJSON)(value.overageCode),
        'profitCode': (0, PackageTransactionCodeType_1.PackageTransactionCodeTypeToJSON)(value.profitCode),
        'lossCode': (0, PackageTransactionCodeType_1.PackageTransactionCodeTypeToJSON)(value.lossCode),
        'alternateCodesList': value.alternateCodesList === undefined ? undefined : (value.alternateCodesList.map(PackageTransactionCodeType_1.PackageTransactionCodeTypeToJSON)),
    };
}
exports.PackagePostingRulesTypeToJSON = PackagePostingRulesTypeToJSON;
