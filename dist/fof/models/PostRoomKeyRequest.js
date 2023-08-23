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
exports.PostRoomKeyRequestToJSON = exports.PostRoomKeyRequestFromJSONTyped = exports.PostRoomKeyRequestFromJSON = exports.instanceOfPostRoomKeyRequest = void 0;
const runtime_1 = require("../runtime");
const InstanceLink_1 = require("./InstanceLink");
const KeyCardType_1 = require("./KeyCardType");
const KeyTrackType_1 = require("./KeyTrackType");
const RoomKeyType_1 = require("./RoomKeyType");
const UniqueIDType_1 = require("./UniqueIDType");
const WarningType_1 = require("./WarningType");
/**
 * Check if a given object implements the PostRoomKeyRequest interface.
 */
function instanceOfPostRoomKeyRequest(value) {
    let isInstance = true;
    return isInstance;
}
exports.instanceOfPostRoomKeyRequest = instanceOfPostRoomKeyRequest;
function PostRoomKeyRequestFromJSON(json) {
    return PostRoomKeyRequestFromJSONTyped(json, false);
}
exports.PostRoomKeyRequestFromJSON = PostRoomKeyRequestFromJSON;
function PostRoomKeyRequestFromJSONTyped(json, ignoreDiscriminator) {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        'reservationIdList': !(0, runtime_1.exists)(json, 'reservationIdList') ? undefined : (json['reservationIdList'].map(UniqueIDType_1.UniqueIDTypeFromJSON)),
        'roomNumber': !(0, runtime_1.exists)(json, 'roomNumber') ? undefined : json['roomNumber'],
        'keyTrack': !(0, runtime_1.exists)(json, 'keyTrack') ? undefined : (0, KeyTrackType_1.KeyTrackTypeFromJSON)(json['keyTrack']),
        'resort': !(0, runtime_1.exists)(json, 'resort') ? undefined : json['resort'],
        'noOfKeys': !(0, runtime_1.exists)(json, 'noOfKeys') ? undefined : json['noOfKeys'],
        'keyValidityStart': !(0, runtime_1.exists)(json, 'keyValidityStart') ? undefined : (new Date(json['keyValidityStart'])),
        'keyValidityEnd': !(0, runtime_1.exists)(json, 'keyValidityEnd') ? undefined : (new Date(json['keyValidityEnd'])),
        'encoderTerminal': !(0, runtime_1.exists)(json, 'encoderTerminal') ? undefined : json['encoderTerminal'],
        'encoderId': !(0, runtime_1.exists)(json, 'encoderId') ? undefined : (0, UniqueIDType_1.UniqueIDTypeFromJSON)(json['encoderId']),
        'keyType': !(0, runtime_1.exists)(json, 'keyType') ? undefined : (0, RoomKeyType_1.RoomKeyTypeFromJSON)(json['keyType']),
        'keyCardType': !(0, runtime_1.exists)(json, 'keyCardType') ? undefined : (0, KeyCardType_1.KeyCardTypeFromJSON)(json['keyCardType']),
        'keyCardUId': !(0, runtime_1.exists)(json, 'keyCardUId') ? undefined : json['keyCardUId'],
        'keyOptions': !(0, runtime_1.exists)(json, 'keyOptions') ? undefined : json['keyOptions'],
        'additionalRooms': !(0, runtime_1.exists)(json, 'additionalRooms') ? undefined : json['additionalRooms'],
        'links': !(0, runtime_1.exists)(json, 'links') ? undefined : (json['links'].map(InstanceLink_1.InstanceLinkFromJSON)),
        'warnings': !(0, runtime_1.exists)(json, 'warnings') ? undefined : (json['warnings'].map(WarningType_1.WarningTypeFromJSON)),
    };
}
exports.PostRoomKeyRequestFromJSONTyped = PostRoomKeyRequestFromJSONTyped;
function PostRoomKeyRequestToJSON(value) {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        'reservationIdList': value.reservationIdList === undefined ? undefined : (value.reservationIdList.map(UniqueIDType_1.UniqueIDTypeToJSON)),
        'roomNumber': value.roomNumber,
        'keyTrack': (0, KeyTrackType_1.KeyTrackTypeToJSON)(value.keyTrack),
        'resort': value.resort,
        'noOfKeys': value.noOfKeys,
        'keyValidityStart': value.keyValidityStart === undefined ? undefined : (value.keyValidityStart.toISOString()),
        'keyValidityEnd': value.keyValidityEnd === undefined ? undefined : (value.keyValidityEnd.toISOString()),
        'encoderTerminal': value.encoderTerminal,
        'encoderId': (0, UniqueIDType_1.UniqueIDTypeToJSON)(value.encoderId),
        'keyType': (0, RoomKeyType_1.RoomKeyTypeToJSON)(value.keyType),
        'keyCardType': (0, KeyCardType_1.KeyCardTypeToJSON)(value.keyCardType),
        'keyCardUId': value.keyCardUId,
        'keyOptions': value.keyOptions,
        'additionalRooms': value.additionalRooms,
        'links': value.links === undefined ? undefined : (value.links.map(InstanceLink_1.InstanceLinkToJSON)),
        'warnings': value.warnings === undefined ? undefined : (value.warnings.map(WarningType_1.WarningTypeToJSON)),
    };
}
exports.PostRoomKeyRequestToJSON = PostRoomKeyRequestToJSON;
