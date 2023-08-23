"use strict";
/* tslint:disable */
/* eslint-disable */
/**
 * OPERA Cloud Room Configuration API
 * APIs to cater for room configuration, such as configuring room types, room Classes, creating new room features, or updating housekeeping room maintenance reasons.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 21.5.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
Object.defineProperty(exports, "__esModule", { value: true });
exports.RoomTypeTemplatesToBeChangedRoomTypeTemplateToJSON = exports.RoomTypeTemplatesToBeChangedRoomTypeTemplateFromJSONTyped = exports.RoomTypeTemplatesToBeChangedRoomTypeTemplateFromJSON = exports.instanceOfRoomTypeTemplatesToBeChangedRoomTypeTemplate = void 0;
const runtime_1 = require("../runtime");
const RoomTypeType_1 = require("./RoomTypeType");
/**
 * Check if a given object implements the RoomTypeTemplatesToBeChangedRoomTypeTemplate interface.
 */
function instanceOfRoomTypeTemplatesToBeChangedRoomTypeTemplate(value) {
    let isInstance = true;
    return isInstance;
}
exports.instanceOfRoomTypeTemplatesToBeChangedRoomTypeTemplate = instanceOfRoomTypeTemplatesToBeChangedRoomTypeTemplate;
function RoomTypeTemplatesToBeChangedRoomTypeTemplateFromJSON(json) {
    return RoomTypeTemplatesToBeChangedRoomTypeTemplateFromJSONTyped(json, false);
}
exports.RoomTypeTemplatesToBeChangedRoomTypeTemplateFromJSON = RoomTypeTemplatesToBeChangedRoomTypeTemplateFromJSON;
function RoomTypeTemplatesToBeChangedRoomTypeTemplateFromJSONTyped(json, ignoreDiscriminator) {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        'roomTypeTemplateDetails': !(0, runtime_1.exists)(json, 'roomTypeTemplateDetails') ? undefined : (0, RoomTypeType_1.RoomTypeTypeFromJSON)(json['roomTypeTemplateDetails']),
    };
}
exports.RoomTypeTemplatesToBeChangedRoomTypeTemplateFromJSONTyped = RoomTypeTemplatesToBeChangedRoomTypeTemplateFromJSONTyped;
function RoomTypeTemplatesToBeChangedRoomTypeTemplateToJSON(value) {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        'roomTypeTemplateDetails': (0, RoomTypeType_1.RoomTypeTypeToJSON)(value.roomTypeTemplateDetails),
    };
}
exports.RoomTypeTemplatesToBeChangedRoomTypeTemplateToJSON = RoomTypeTemplatesToBeChangedRoomTypeTemplateToJSON;
