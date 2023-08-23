"use strict";
/* tslint:disable */
/* eslint-disable */
/**
 * OPERA Cloud Integration Configuration API
 * APIs catering to Integration Configuration in OPERA Cloud.  Operations such as get Hotel Interface Types, or get UDF mappings can be found in this module.<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 22.3.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
Object.defineProperty(exports, "__esModule", { value: true });
exports.ImportInterfaceRoomsTypeToJSON = exports.ImportInterfaceRoomsTypeFromJSONTyped = exports.ImportInterfaceRoomsTypeFromJSON = exports.instanceOfImportInterfaceRoomsType = void 0;
const runtime_1 = require("../runtime");
const DataLineType_1 = require("./DataLineType");
const RoomMaskSetupType_1 = require("./RoomMaskSetupType");
/**
 * Check if a given object implements the ImportInterfaceRoomsType interface.
 */
function instanceOfImportInterfaceRoomsType(value) {
    let isInstance = true;
    return isInstance;
}
exports.instanceOfImportInterfaceRoomsType = instanceOfImportInterfaceRoomsType;
function ImportInterfaceRoomsTypeFromJSON(json) {
    return ImportInterfaceRoomsTypeFromJSONTyped(json, false);
}
exports.ImportInterfaceRoomsTypeFromJSON = ImportInterfaceRoomsTypeFromJSON;
function ImportInterfaceRoomsTypeFromJSONTyped(json, ignoreDiscriminator) {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        'hotelId': !(0, runtime_1.exists)(json, 'hotelId') ? undefined : json['hotelId'],
        'logo': !(0, runtime_1.exists)(json, 'logo') ? undefined : json['logo'],
        'dataLine': !(0, runtime_1.exists)(json, 'dataLine') ? undefined : (0, DataLineType_1.DataLineTypeFromJSON)(json['dataLine']),
        'numberOfLines': !(0, runtime_1.exists)(json, 'numberOfLines') ? undefined : json['numberOfLines'],
        'removeLeadingZeros': !(0, runtime_1.exists)(json, 'removeLeadingZeros') ? undefined : json['removeLeadingZeros'],
        'line1Prefix': !(0, runtime_1.exists)(json, 'line1Prefix') ? undefined : json['line1Prefix'],
        'line1AddTo': !(0, runtime_1.exists)(json, 'line1AddTo') ? undefined : json['line1AddTo'],
        'line2Prefix': !(0, runtime_1.exists)(json, 'line2Prefix') ? undefined : json['line2Prefix'],
        'line2AddTo': !(0, runtime_1.exists)(json, 'line2AddTo') ? undefined : json['line2AddTo'],
        'line3Prefix': !(0, runtime_1.exists)(json, 'line3Prefix') ? undefined : json['line3Prefix'],
        'line3AddTo': !(0, runtime_1.exists)(json, 'line3AddTo') ? undefined : json['line3AddTo'],
        'line4Prefix': !(0, runtime_1.exists)(json, 'line4Prefix') ? undefined : json['line4Prefix'],
        'line4AddTo': !(0, runtime_1.exists)(json, 'line4AddTo') ? undefined : json['line4AddTo'],
        'selectedRoomTypes': !(0, runtime_1.exists)(json, 'selectedRoomTypes') ? undefined : json['selectedRoomTypes'],
        'maskLines': !(0, runtime_1.exists)(json, 'maskLines') ? undefined : (json['maskLines'].map(RoomMaskSetupType_1.RoomMaskSetupTypeFromJSON)),
    };
}
exports.ImportInterfaceRoomsTypeFromJSONTyped = ImportInterfaceRoomsTypeFromJSONTyped;
function ImportInterfaceRoomsTypeToJSON(value) {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        'hotelId': value.hotelId,
        'logo': value.logo,
        'dataLine': (0, DataLineType_1.DataLineTypeToJSON)(value.dataLine),
        'numberOfLines': value.numberOfLines,
        'removeLeadingZeros': value.removeLeadingZeros,
        'line1Prefix': value.line1Prefix,
        'line1AddTo': value.line1AddTo,
        'line2Prefix': value.line2Prefix,
        'line2AddTo': value.line2AddTo,
        'line3Prefix': value.line3Prefix,
        'line3AddTo': value.line3AddTo,
        'line4Prefix': value.line4Prefix,
        'line4AddTo': value.line4AddTo,
        'selectedRoomTypes': value.selectedRoomTypes,
        'maskLines': value.maskLines === undefined ? undefined : (value.maskLines.map(RoomMaskSetupType_1.RoomMaskSetupTypeToJSON)),
    };
}
exports.ImportInterfaceRoomsTypeToJSON = ImportInterfaceRoomsTypeToJSON;
