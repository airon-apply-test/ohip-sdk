"use strict";
/* tslint:disable */
/* eslint-disable */
/**
 * OPERA Cloud Front Desk Operations Service
 * APIs to cater for Front Desk Operations and Front Desk Statistic functionality in OPERA Cloud. <br /><br /> Front Desk features some of the most commonly used operations in OPERA Cloud, such as managing guest arrivals, managing in-house guests, and managing guest departures. Some additional tasks you can complete from the Front Desk menu are room searches, room assignments, and quick check outs as well as opening folios, creating registration cards, setting wake up calls, and sending messages to guests.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 21.5.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
Object.defineProperty(exports, "__esModule", { value: true });
exports.CodeInfoTypeToJSON = exports.CodeInfoTypeFromJSONTyped = exports.CodeInfoTypeFromJSON = exports.instanceOfCodeInfoType = void 0;
const runtime_1 = require("../runtime");
const AddtionalCodeInfoTypeInner_1 = require("./AddtionalCodeInfoTypeInner");
/**
 * Check if a given object implements the CodeInfoType interface.
 */
function instanceOfCodeInfoType(value) {
    let isInstance = true;
    return isInstance;
}
exports.instanceOfCodeInfoType = instanceOfCodeInfoType;
function CodeInfoTypeFromJSON(json) {
    return CodeInfoTypeFromJSONTyped(json, false);
}
exports.CodeInfoTypeFromJSON = CodeInfoTypeFromJSON;
function CodeInfoTypeFromJSONTyped(json, ignoreDiscriminator) {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        'description': !(0, runtime_1.exists)(json, 'description') ? undefined : json['description'],
        'addtionalCodeInfo': !(0, runtime_1.exists)(json, 'addtionalCodeInfo') ? undefined : (json['addtionalCodeInfo'].map(AddtionalCodeInfoTypeInner_1.AddtionalCodeInfoTypeInnerFromJSON)),
        'hotelId': !(0, runtime_1.exists)(json, 'hotelId') ? undefined : json['hotelId'],
        'code': !(0, runtime_1.exists)(json, 'code') ? undefined : json['code'],
    };
}
exports.CodeInfoTypeFromJSONTyped = CodeInfoTypeFromJSONTyped;
function CodeInfoTypeToJSON(value) {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        'description': value.description,
        'addtionalCodeInfo': value.addtionalCodeInfo === undefined ? undefined : (value.addtionalCodeInfo.map(AddtionalCodeInfoTypeInner_1.AddtionalCodeInfoTypeInnerToJSON)),
        'hotelId': value.hotelId,
        'code': value.code,
    };
}
exports.CodeInfoTypeToJSON = CodeInfoTypeToJSON;
