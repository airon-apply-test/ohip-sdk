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
exports.ConfigPackageTransactionTypeToJSON = exports.ConfigPackageTransactionTypeFromJSONTyped = exports.ConfigPackageTransactionTypeFromJSON = exports.instanceOfConfigPackageTransactionType = void 0;
const runtime_1 = require("../runtime");
const PackagePostingRulesType_1 = require("./PackagePostingRulesType");
/**
 * Check if a given object implements the ConfigPackageTransactionType interface.
 */
function instanceOfConfigPackageTransactionType(value) {
    let isInstance = true;
    return isInstance;
}
exports.instanceOfConfigPackageTransactionType = instanceOfConfigPackageTransactionType;
function ConfigPackageTransactionTypeFromJSON(json) {
    return ConfigPackageTransactionTypeFromJSONTyped(json, false);
}
exports.ConfigPackageTransactionTypeFromJSON = ConfigPackageTransactionTypeFromJSON;
function ConfigPackageTransactionTypeFromJSONTyped(json, ignoreDiscriminator) {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        'allowance': !(0, runtime_1.exists)(json, 'allowance') ? undefined : json['allowance'],
        'currency': !(0, runtime_1.exists)(json, 'currency') ? undefined : json['currency'],
        'postingType': !(0, runtime_1.exists)(json, 'postingType') ? undefined : json['postingType'],
        'calculationRule': !(0, runtime_1.exists)(json, 'calculationRule') ? undefined : json['calculationRule'],
        'packagePostingRules': !(0, runtime_1.exists)(json, 'packagePostingRules') ? undefined : (0, PackagePostingRulesType_1.PackagePostingRulesTypeFromJSON)(json['packagePostingRules']),
    };
}
exports.ConfigPackageTransactionTypeFromJSONTyped = ConfigPackageTransactionTypeFromJSONTyped;
function ConfigPackageTransactionTypeToJSON(value) {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        'allowance': value.allowance,
        'currency': value.currency,
        'postingType': value.postingType,
        'calculationRule': value.calculationRule,
        'packagePostingRules': (0, PackagePostingRulesType_1.PackagePostingRulesTypeToJSON)(value.packagePostingRules),
    };
}
exports.ConfigPackageTransactionTypeToJSON = ConfigPackageTransactionTypeToJSON;
