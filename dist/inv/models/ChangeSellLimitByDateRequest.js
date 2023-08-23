"use strict";
/* tslint:disable */
/* eslint-disable */
/**
 * OPERA Cloud Inventory API
 * APIs to cater for Inventory functionality in OPERA Cloud. This includes sell limits for date ranges, viewing and updating the property&apos;s inventory, as well as item inventory (such as rollaways, microwaves etc.).<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 22.3.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
Object.defineProperty(exports, "__esModule", { value: true });
exports.ChangeSellLimitByDateRequestToJSON = exports.ChangeSellLimitByDateRequestFromJSONTyped = exports.ChangeSellLimitByDateRequestFromJSON = exports.instanceOfChangeSellLimitByDateRequest = void 0;
const runtime_1 = require("../runtime");
const SellLimitByDateType_1 = require("./SellLimitByDateType");
const WarningType_1 = require("./WarningType");
/**
 * Check if a given object implements the ChangeSellLimitByDateRequest interface.
 */
function instanceOfChangeSellLimitByDateRequest(value) {
    let isInstance = true;
    return isInstance;
}
exports.instanceOfChangeSellLimitByDateRequest = instanceOfChangeSellLimitByDateRequest;
function ChangeSellLimitByDateRequestFromJSON(json) {
    return ChangeSellLimitByDateRequestFromJSONTyped(json, false);
}
exports.ChangeSellLimitByDateRequestFromJSON = ChangeSellLimitByDateRequestFromJSON;
function ChangeSellLimitByDateRequestFromJSONTyped(json, ignoreDiscriminator) {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        'sellLimitsByDate': !(0, runtime_1.exists)(json, 'sellLimitsByDate') ? undefined : (0, SellLimitByDateType_1.SellLimitByDateTypeFromJSON)(json['sellLimitsByDate']),
        'warnings': !(0, runtime_1.exists)(json, 'warnings') ? undefined : (json['warnings'].map(WarningType_1.WarningTypeFromJSON)),
    };
}
exports.ChangeSellLimitByDateRequestFromJSONTyped = ChangeSellLimitByDateRequestFromJSONTyped;
function ChangeSellLimitByDateRequestToJSON(value) {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        'sellLimitsByDate': (0, SellLimitByDateType_1.SellLimitByDateTypeToJSON)(value.sellLimitsByDate),
        'warnings': value.warnings === undefined ? undefined : (value.warnings.map(WarningType_1.WarningTypeToJSON)),
    };
}
exports.ChangeSellLimitByDateRequestToJSON = ChangeSellLimitByDateRequestToJSON;
