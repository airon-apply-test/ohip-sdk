"use strict";
/* tslint:disable */
/* eslint-disable */
/**
 * OPERA Cloud CRM Configuration API
 * APIs for Customer Relationship Management (profile) configuration, such as creating preferences, or address types.  It also includes Membership Configuration, where you can retrieve membership levels that are configured for a property, or create new membership enrollment codes.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 21.5.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
Object.defineProperty(exports, "__esModule", { value: true });
exports.VIPLevelTypeToJSON = exports.VIPLevelTypeFromJSONTyped = exports.VIPLevelTypeFromJSON = exports.instanceOfVIPLevelType = void 0;
const runtime_1 = require("../runtime");
const CommonMasterColorType_1 = require("./CommonMasterColorType");
const TranslationTextType2000_1 = require("./TranslationTextType2000");
/**
 * Check if a given object implements the VIPLevelType interface.
 */
function instanceOfVIPLevelType(value) {
    let isInstance = true;
    return isInstance;
}
exports.instanceOfVIPLevelType = instanceOfVIPLevelType;
function VIPLevelTypeFromJSON(json) {
    return VIPLevelTypeFromJSONTyped(json, false);
}
exports.VIPLevelTypeFromJSON = VIPLevelTypeFromJSON;
function VIPLevelTypeFromJSONTyped(json, ignoreDiscriminator) {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        'code': !(0, runtime_1.exists)(json, 'code') ? undefined : json['code'],
        'description': !(0, runtime_1.exists)(json, 'description') ? undefined : (0, TranslationTextType2000_1.TranslationTextType2000FromJSON)(json['description']),
        'displayOrder': !(0, runtime_1.exists)(json, 'displayOrder') ? undefined : json['displayOrder'],
        'displayColor': !(0, runtime_1.exists)(json, 'displayColor') ? undefined : (0, CommonMasterColorType_1.CommonMasterColorTypeFromJSON)(json['displayColor']),
        'ranking': !(0, runtime_1.exists)(json, 'ranking') ? undefined : json['ranking'],
    };
}
exports.VIPLevelTypeFromJSONTyped = VIPLevelTypeFromJSONTyped;
function VIPLevelTypeToJSON(value) {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        'code': value.code,
        'description': (0, TranslationTextType2000_1.TranslationTextType2000ToJSON)(value.description),
        'displayOrder': value.displayOrder,
        'displayColor': (0, CommonMasterColorType_1.CommonMasterColorTypeToJSON)(value.displayColor),
        'ranking': value.ranking,
    };
}
exports.VIPLevelTypeToJSON = VIPLevelTypeToJSON;
