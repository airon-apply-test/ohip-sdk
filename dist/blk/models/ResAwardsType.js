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
exports.ResAwardsTypeToJSON = exports.ResAwardsTypeFromJSONTyped = exports.ResAwardsTypeFromJSON = exports.instanceOfResAwardsType = void 0;
const runtime_1 = require("../runtime");
const AwardVouchersTypeInner_1 = require("./AwardVouchersTypeInner");
const UniqueIDType_1 = require("./UniqueIDType");
/**
 * Check if a given object implements the ResAwardsType interface.
 */
function instanceOfResAwardsType(value) {
    let isInstance = true;
    return isInstance;
}
exports.instanceOfResAwardsType = instanceOfResAwardsType;
function ResAwardsTypeFromJSON(json) {
    return ResAwardsTypeFromJSONTyped(json, false);
}
exports.ResAwardsTypeFromJSON = ResAwardsTypeFromJSON;
function ResAwardsTypeFromJSONTyped(json, ignoreDiscriminator) {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        'membershipNo': !(0, runtime_1.exists)(json, 'membershipNo') ? undefined : (0, UniqueIDType_1.UniqueIDTypeFromJSON)(json['membershipNo']),
        'awardVouchers': !(0, runtime_1.exists)(json, 'awardVouchers') ? undefined : (json['awardVouchers'].map(AwardVouchersTypeInner_1.AwardVouchersTypeInnerFromJSON)),
        'originalRoomType': !(0, runtime_1.exists)(json, 'originalRoomType') ? undefined : json['originalRoomType'],
        'upgradeRoomType': !(0, runtime_1.exists)(json, 'upgradeRoomType') ? undefined : json['upgradeRoomType'],
    };
}
exports.ResAwardsTypeFromJSONTyped = ResAwardsTypeFromJSONTyped;
function ResAwardsTypeToJSON(value) {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        'membershipNo': (0, UniqueIDType_1.UniqueIDTypeToJSON)(value.membershipNo),
        'awardVouchers': value.awardVouchers === undefined ? undefined : (value.awardVouchers.map(AwardVouchersTypeInner_1.AwardVouchersTypeInnerToJSON)),
        'originalRoomType': value.originalRoomType,
        'upgradeRoomType': value.upgradeRoomType,
    };
}
exports.ResAwardsTypeToJSON = ResAwardsTypeToJSON;
