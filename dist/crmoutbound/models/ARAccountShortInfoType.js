"use strict";
/* tslint:disable */
/* eslint-disable */
/**
 * OPERA Cloud Customer Relationship Management Outbound API
 * APIs to cater for Customer Relationship Management external (outbound) functionality with OPERA. These APIs facilitate various operations related to getting data from an external system, and inserting it into OPERA.<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 22.3.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
Object.defineProperty(exports, "__esModule", { value: true });
exports.ARAccountShortInfoTypeToJSON = exports.ARAccountShortInfoTypeFromJSONTyped = exports.ARAccountShortInfoTypeFromJSON = exports.instanceOfARAccountShortInfoType = void 0;
const runtime_1 = require("../runtime");
const ARAccountStatusType_1 = require("./ARAccountStatusType");
const UniqueIDType_1 = require("./UniqueIDType");
/**
 * Check if a given object implements the ARAccountShortInfoType interface.
 */
function instanceOfARAccountShortInfoType(value) {
    let isInstance = true;
    return isInstance;
}
exports.instanceOfARAccountShortInfoType = instanceOfARAccountShortInfoType;
function ARAccountShortInfoTypeFromJSON(json) {
    return ARAccountShortInfoTypeFromJSONTyped(json, false);
}
exports.ARAccountShortInfoTypeFromJSON = ARAccountShortInfoTypeFromJSON;
function ARAccountShortInfoTypeFromJSONTyped(json, ignoreDiscriminator) {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        'accountName': !(0, runtime_1.exists)(json, 'accountName') ? undefined : json['accountName'],
        'accountId': !(0, runtime_1.exists)(json, 'accountId') ? undefined : (0, UniqueIDType_1.UniqueIDTypeFromJSON)(json['accountId']),
        'accountNo': !(0, runtime_1.exists)(json, 'accountNo') ? undefined : json['accountNo'],
        'status': !(0, runtime_1.exists)(json, 'status') ? undefined : (0, ARAccountStatusType_1.ARAccountStatusTypeFromJSON)(json['status']),
    };
}
exports.ARAccountShortInfoTypeFromJSONTyped = ARAccountShortInfoTypeFromJSONTyped;
function ARAccountShortInfoTypeToJSON(value) {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        'accountName': value.accountName,
        'accountId': (0, UniqueIDType_1.UniqueIDTypeToJSON)(value.accountId),
        'accountNo': value.accountNo,
        'status': (0, ARAccountStatusType_1.ARAccountStatusTypeToJSON)(value.status),
    };
}
exports.ARAccountShortInfoTypeToJSON = ARAccountShortInfoTypeToJSON;
