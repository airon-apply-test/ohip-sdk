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
exports.CurrentRoomInfoTypeToJSON = exports.CurrentRoomInfoTypeFromJSONTyped = exports.CurrentRoomInfoTypeFromJSON = exports.instanceOfCurrentRoomInfoType = void 0;
const runtime_1 = require("../runtime");
/**
 * Check if a given object implements the CurrentRoomInfoType interface.
 */
function instanceOfCurrentRoomInfoType(value) {
    let isInstance = true;
    return isInstance;
}
exports.instanceOfCurrentRoomInfoType = instanceOfCurrentRoomInfoType;
function CurrentRoomInfoTypeFromJSON(json) {
    return CurrentRoomInfoTypeFromJSONTyped(json, false);
}
exports.CurrentRoomInfoTypeFromJSON = CurrentRoomInfoTypeFromJSON;
function CurrentRoomInfoTypeFromJSONTyped(json, ignoreDiscriminator) {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        'roomType': !(0, runtime_1.exists)(json, 'roomType') ? undefined : json['roomType'],
        'roomId': !(0, runtime_1.exists)(json, 'roomId') ? undefined : json['roomId'],
        'suggestedRoomNumbers': !(0, runtime_1.exists)(json, 'suggestedRoomNumbers') ? undefined : json['suggestedRoomNumbers'],
        'roomDescription': !(0, runtime_1.exists)(json, 'roomDescription') ? undefined : json['roomDescription'],
        'roomViewCode': !(0, runtime_1.exists)(json, 'roomViewCode') ? undefined : json['roomViewCode'],
        'assignedByAI': !(0, runtime_1.exists)(json, 'assignedByAI') ? undefined : json['assignedByAI'],
        'upgradedByAI': !(0, runtime_1.exists)(json, 'upgradedByAI') ? undefined : json['upgradedByAI'],
    };
}
exports.CurrentRoomInfoTypeFromJSONTyped = CurrentRoomInfoTypeFromJSONTyped;
function CurrentRoomInfoTypeToJSON(value) {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        'roomType': value.roomType,
        'roomId': value.roomId,
        'suggestedRoomNumbers': value.suggestedRoomNumbers,
        'roomDescription': value.roomDescription,
        'roomViewCode': value.roomViewCode,
        'assignedByAI': value.assignedByAI,
        'upgradedByAI': value.upgradedByAI,
    };
}
exports.CurrentRoomInfoTypeToJSON = CurrentRoomInfoTypeToJSON;
