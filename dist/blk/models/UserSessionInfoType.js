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
exports.UserSessionInfoTypeToJSON = exports.UserSessionInfoTypeFromJSONTyped = exports.UserSessionInfoTypeFromJSON = exports.instanceOfUserSessionInfoType = void 0;
const runtime_1 = require("../runtime");
const CodeDescriptionType_1 = require("./CodeDescriptionType");
const ParameterType_1 = require("./ParameterType");
const UserSessionDefaultsType_1 = require("./UserSessionDefaultsType");
/**
 * Check if a given object implements the UserSessionInfoType interface.
 */
function instanceOfUserSessionInfoType(value) {
    let isInstance = true;
    return isInstance;
}
exports.instanceOfUserSessionInfoType = instanceOfUserSessionInfoType;
function UserSessionInfoTypeFromJSON(json) {
    return UserSessionInfoTypeFromJSONTyped(json, false);
}
exports.UserSessionInfoTypeFromJSON = UserSessionInfoTypeFromJSON;
function UserSessionInfoTypeFromJSONTyped(json, ignoreDiscriminator) {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        'businessDate': !(0, runtime_1.exists)(json, 'businessDate') ? undefined : (new Date(json['businessDate'])),
        'systemDate': !(0, runtime_1.exists)(json, 'systemDate') ? undefined : (new Date(json['systemDate'])),
        'terminal': !(0, runtime_1.exists)(json, 'terminal') ? undefined : json['terminal'],
        'runningApp': !(0, runtime_1.exists)(json, 'runningApp') ? undefined : json['runningApp'],
        'shareProfiles': !(0, runtime_1.exists)(json, 'shareProfiles') ? undefined : json['shareProfiles'],
        'hotel': !(0, runtime_1.exists)(json, 'hotel') ? undefined : (0, CodeDescriptionType_1.CodeDescriptionTypeFromJSON)(json['hotel']),
        'cro': !(0, runtime_1.exists)(json, 'cro') ? undefined : (0, CodeDescriptionType_1.CodeDescriptionTypeFromJSON)(json['cro']),
        'chain': !(0, runtime_1.exists)(json, 'chain') ? undefined : json['chain'],
        'cROCountryCode': !(0, runtime_1.exists)(json, 'cROCountryCode') ? undefined : json['cROCountryCode'],
        'sessionDefaults': !(0, runtime_1.exists)(json, 'sessionDefaults') ? undefined : (0, UserSessionDefaultsType_1.UserSessionDefaultsTypeFromJSON)(json['sessionDefaults']),
        'parameters': !(0, runtime_1.exists)(json, 'parameters') ? undefined : (json['parameters'].map(ParameterType_1.ParameterTypeFromJSON)),
    };
}
exports.UserSessionInfoTypeFromJSONTyped = UserSessionInfoTypeFromJSONTyped;
function UserSessionInfoTypeToJSON(value) {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        'businessDate': value.businessDate === undefined ? undefined : (value.businessDate.toISOString().substr(0, 10)),
        'systemDate': value.systemDate === undefined ? undefined : (value.systemDate.toISOString().substr(0, 10)),
        'terminal': value.terminal,
        'runningApp': value.runningApp,
        'shareProfiles': value.shareProfiles,
        'hotel': (0, CodeDescriptionType_1.CodeDescriptionTypeToJSON)(value.hotel),
        'cro': (0, CodeDescriptionType_1.CodeDescriptionTypeToJSON)(value.cro),
        'chain': value.chain,
        'cROCountryCode': value.cROCountryCode,
        'sessionDefaults': (0, UserSessionDefaultsType_1.UserSessionDefaultsTypeToJSON)(value.sessionDefaults),
        'parameters': value.parameters === undefined ? undefined : (value.parameters.map(ParameterType_1.ParameterTypeToJSON)),
    };
}
exports.UserSessionInfoTypeToJSON = UserSessionInfoTypeToJSON;
