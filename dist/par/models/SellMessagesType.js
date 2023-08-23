"use strict";
/* tslint:disable */
/* eslint-disable */
/**
 * OPERA Cloud Price Availability Rate API
 * APIs to cater for Price and Rate Availability functionality in OPERA Cloud. <br /><br />Availability enables you to manage your room inventory by providing a detailed view of all available and sold rooms at a property. Some of the tasks you can perform include defining conditions for stay restrictions, setting room sell limits, and searching for and viewing room availability.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 21.5.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
Object.defineProperty(exports, "__esModule", { value: true });
exports.SellMessagesTypeToJSON = exports.SellMessagesTypeFromJSONTyped = exports.SellMessagesTypeFromJSON = exports.instanceOfSellMessagesType = void 0;
const runtime_1 = require("../runtime");
const SellMessageType_1 = require("./SellMessageType");
/**
 * Check if a given object implements the SellMessagesType interface.
 */
function instanceOfSellMessagesType(value) {
    let isInstance = true;
    return isInstance;
}
exports.instanceOfSellMessagesType = instanceOfSellMessagesType;
function SellMessagesTypeFromJSON(json) {
    return SellMessagesTypeFromJSONTyped(json, false);
}
exports.SellMessagesTypeFromJSON = SellMessagesTypeFromJSON;
function SellMessagesTypeFromJSONTyped(json, ignoreDiscriminator) {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        'sellMessage': !(0, runtime_1.exists)(json, 'sellMessage') ? undefined : (json['sellMessage'].map(SellMessageType_1.SellMessageTypeFromJSON)),
        'hasMore': !(0, runtime_1.exists)(json, 'hasMore') ? undefined : json['hasMore'],
        'totalResults': !(0, runtime_1.exists)(json, 'totalResults') ? undefined : json['totalResults'],
        'count': !(0, runtime_1.exists)(json, 'count') ? undefined : json['count'],
    };
}
exports.SellMessagesTypeFromJSONTyped = SellMessagesTypeFromJSONTyped;
function SellMessagesTypeToJSON(value) {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        'sellMessage': value.sellMessage === undefined ? undefined : (value.sellMessage.map(SellMessageType_1.SellMessageTypeToJSON)),
        'hasMore': value.hasMore,
        'totalResults': value.totalResults,
        'count': value.count,
    };
}
exports.SellMessagesTypeToJSON = SellMessagesTypeToJSON;
