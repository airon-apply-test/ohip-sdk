"use strict";
/* tslint:disable */
/* eslint-disable */
/**
 * OPERA Cloud Block Configuration API
 * APIs for Block configuration, such as creating, updating, fetching and removing codes related to blocks. <br />< This might include fetching the block cancellation reasons, or creating new block refused reasons.  Wash schedules can be create, or new reservation methods could be added for a property.<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 22.3.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
Object.defineProperty(exports, "__esModule", { value: true });
exports.BlockStatusCodeTypeToJSON = exports.BlockStatusCodeTypeFromJSONTyped = exports.BlockStatusCodeTypeFromJSON = exports.instanceOfBlockStatusCodeType = void 0;
const runtime_1 = require("../runtime");
const BlockStatusCodeStatusType_1 = require("./BlockStatusCodeStatusType");
const CodeDescriptionType_1 = require("./CodeDescriptionType");
const StatusColorType_1 = require("./StatusColorType");
/**
 * Check if a given object implements the BlockStatusCodeType interface.
 */
function instanceOfBlockStatusCodeType(value) {
    let isInstance = true;
    return isInstance;
}
exports.instanceOfBlockStatusCodeType = instanceOfBlockStatusCodeType;
function BlockStatusCodeTypeFromJSON(json) {
    return BlockStatusCodeTypeFromJSONTyped(json, false);
}
exports.BlockStatusCodeTypeFromJSON = BlockStatusCodeTypeFromJSON;
function BlockStatusCodeTypeFromJSONTyped(json, ignoreDiscriminator) {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        'status': !(0, runtime_1.exists)(json, 'status') ? undefined : (0, CodeDescriptionType_1.CodeDescriptionTypeFromJSON)(json['status']),
        'roomStatusType': !(0, runtime_1.exists)(json, 'roomStatusType') ? undefined : (0, BlockStatusCodeStatusType_1.BlockStatusCodeStatusTypeFromJSON)(json['roomStatusType']),
        'cateringStatusType': !(0, runtime_1.exists)(json, 'cateringStatusType') ? undefined : (0, BlockStatusCodeStatusType_1.BlockStatusCodeStatusTypeFromJSON)(json['cateringStatusType']),
        'defaultReservationType': !(0, runtime_1.exists)(json, 'defaultReservationType') ? undefined : (0, CodeDescriptionType_1.CodeDescriptionTypeFromJSON)(json['defaultReservationType']),
        'reasonType': !(0, runtime_1.exists)(json, 'reasonType') ? undefined : json['reasonType'],
        'color': !(0, runtime_1.exists)(json, 'color') ? undefined : (0, StatusColorType_1.StatusColorTypeFromJSON)(json['color']),
        'allowPickup': !(0, runtime_1.exists)(json, 'allowPickup') ? undefined : json['allowPickup'],
        'returnToInventory': !(0, runtime_1.exists)(json, 'returnToInventory') ? undefined : json['returnToInventory'],
        'starting': !(0, runtime_1.exists)(json, 'starting') ? undefined : json['starting'],
        'leadStatus': !(0, runtime_1.exists)(json, 'leadStatus') ? undefined : json['leadStatus'],
        'logCatering': !(0, runtime_1.exists)(json, 'logCatering') ? undefined : json['logCatering'],
        'showDiary': !(0, runtime_1.exists)(json, 'showDiary') ? undefined : json['showDiary'],
        'orderBy': !(0, runtime_1.exists)(json, 'orderBy') ? undefined : json['orderBy'],
        'inUse': !(0, runtime_1.exists)(json, 'inUse') ? undefined : json['inUse'],
        'cateringInUse': !(0, runtime_1.exists)(json, 'cateringInUse') ? undefined : json['cateringInUse'],
        'opportunity': !(0, runtime_1.exists)(json, 'opportunity') ? undefined : json['opportunity'],
    };
}
exports.BlockStatusCodeTypeFromJSONTyped = BlockStatusCodeTypeFromJSONTyped;
function BlockStatusCodeTypeToJSON(value) {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        'status': (0, CodeDescriptionType_1.CodeDescriptionTypeToJSON)(value.status),
        'roomStatusType': (0, BlockStatusCodeStatusType_1.BlockStatusCodeStatusTypeToJSON)(value.roomStatusType),
        'cateringStatusType': (0, BlockStatusCodeStatusType_1.BlockStatusCodeStatusTypeToJSON)(value.cateringStatusType),
        'defaultReservationType': (0, CodeDescriptionType_1.CodeDescriptionTypeToJSON)(value.defaultReservationType),
        'reasonType': value.reasonType,
        'color': (0, StatusColorType_1.StatusColorTypeToJSON)(value.color),
        'allowPickup': value.allowPickup,
        'returnToInventory': value.returnToInventory,
        'starting': value.starting,
        'leadStatus': value.leadStatus,
        'logCatering': value.logCatering,
        'showDiary': value.showDiary,
        'orderBy': value.orderBy,
        'inUse': value.inUse,
        'cateringInUse': value.cateringInUse,
        'opportunity': value.opportunity,
    };
}
exports.BlockStatusCodeTypeToJSON = BlockStatusCodeTypeToJSON;
