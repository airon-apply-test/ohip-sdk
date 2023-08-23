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
exports.CompAccountingTypeToJSON = exports.CompAccountingTypeFromJSONTyped = exports.CompAccountingTypeFromJSON = exports.instanceOfCompAccountingType = void 0;
const runtime_1 = require("../runtime");
const DateRangeType_1 = require("./DateRangeType");
const UniqueIDType_1 = require("./UniqueIDType");
/**
 * Check if a given object implements the CompAccountingType interface.
 */
function instanceOfCompAccountingType(value) {
    let isInstance = true;
    return isInstance;
}
exports.instanceOfCompAccountingType = instanceOfCompAccountingType;
function CompAccountingTypeFromJSON(json) {
    return CompAccountingTypeFromJSONTyped(json, false);
}
exports.CompAccountingTypeFromJSON = CompAccountingTypeFromJSON;
function CompAccountingTypeFromJSONTyped(json, ignoreDiscriminator) {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        'authorizerId': !(0, runtime_1.exists)(json, 'authorizerId') ? undefined : (0, UniqueIDType_1.UniqueIDTypeFromJSON)(json['authorizerId']),
        'authorizerUserName': !(0, runtime_1.exists)(json, 'authorizerUserName') ? undefined : json['authorizerUserName'],
        'authorizerName': !(0, runtime_1.exists)(json, 'authorizerName') ? undefined : json['authorizerName'],
        'authorizerRateCode': !(0, runtime_1.exists)(json, 'authorizerRateCode') ? undefined : json['authorizerRateCode'],
        'inheritAuthorizerRateCode': !(0, runtime_1.exists)(json, 'inheritAuthorizerRateCode') ? undefined : json['inheritAuthorizerRateCode'],
        'aRNumber': !(0, runtime_1.exists)(json, 'aRNumber') ? undefined : json['aRNumber'],
        'compTypeCode': !(0, runtime_1.exists)(json, 'compTypeCode') ? undefined : json['compTypeCode'],
        'compVoucherNo': !(0, runtime_1.exists)(json, 'compVoucherNo') ? undefined : json['compVoucherNo'],
        'dateRange': !(0, runtime_1.exists)(json, 'dateRange') ? undefined : (0, DateRangeType_1.DateRangeTypeFromJSON)(json['dateRange']),
        'membershipId': !(0, runtime_1.exists)(json, 'membershipId') ? undefined : json['membershipId'],
        'membershipIdNo': !(0, runtime_1.exists)(json, 'membershipIdNo') ? undefined : json['membershipIdNo'],
        'membershipNameOnCard': !(0, runtime_1.exists)(json, 'membershipNameOnCard') ? undefined : json['membershipNameOnCard'],
    };
}
exports.CompAccountingTypeFromJSONTyped = CompAccountingTypeFromJSONTyped;
function CompAccountingTypeToJSON(value) {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        'authorizerId': (0, UniqueIDType_1.UniqueIDTypeToJSON)(value.authorizerId),
        'authorizerUserName': value.authorizerUserName,
        'authorizerName': value.authorizerName,
        'authorizerRateCode': value.authorizerRateCode,
        'inheritAuthorizerRateCode': value.inheritAuthorizerRateCode,
        'aRNumber': value.aRNumber,
        'compTypeCode': value.compTypeCode,
        'compVoucherNo': value.compVoucherNo,
        'dateRange': (0, DateRangeType_1.DateRangeTypeToJSON)(value.dateRange),
        'membershipId': value.membershipId,
        'membershipIdNo': value.membershipIdNo,
        'membershipNameOnCard': value.membershipNameOnCard,
    };
}
exports.CompAccountingTypeToJSON = CompAccountingTypeToJSON;
