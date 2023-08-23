"use strict";
/* tslint:disable */
/* eslint-disable */
/**
 * OPERA Cloud Enterprise Configuration API
 * APIs to cater for Enterprise Configuration functionality in OPERA Cloud. <br /><br In this module, you can configure a variety of options related to your properties such as their locations, facilities, and local attractions. The available options are dependant on the active controls at your property.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 21.5.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
Object.defineProperty(exports, "__esModule", { value: true });
exports.WebUserAccountTypeToJSON = exports.WebUserAccountTypeFromJSONTyped = exports.WebUserAccountTypeFromJSON = exports.instanceOfWebUserAccountType = void 0;
const runtime_1 = require("../runtime");
const CodeDescriptionType_1 = require("./CodeDescriptionType");
const WebUserHistoryType_1 = require("./WebUserHistoryType");
/**
 * Check if a given object implements the WebUserAccountType interface.
 */
function instanceOfWebUserAccountType(value) {
    let isInstance = true;
    return isInstance;
}
exports.instanceOfWebUserAccountType = instanceOfWebUserAccountType;
function WebUserAccountTypeFromJSON(json) {
    return WebUserAccountTypeFromJSONTyped(json, false);
}
exports.WebUserAccountTypeFromJSON = WebUserAccountTypeFromJSON;
function WebUserAccountTypeFromJSONTyped(json, ignoreDiscriminator) {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        'securityQuestion': !(0, runtime_1.exists)(json, 'securityQuestion') ? undefined : (0, CodeDescriptionType_1.CodeDescriptionTypeFromJSON)(json['securityQuestion']),
        'securityAnswer': !(0, runtime_1.exists)(json, 'securityAnswer') ? undefined : json['securityAnswer'],
        'comments': !(0, runtime_1.exists)(json, 'comments') ? undefined : json['comments'],
        'history': !(0, runtime_1.exists)(json, 'history') ? undefined : (0, WebUserHistoryType_1.WebUserHistoryTypeFromJSON)(json['history']),
        'newLoginName': !(0, runtime_1.exists)(json, 'newLoginName') ? undefined : json['newLoginName'],
        'newPassword': !(0, runtime_1.exists)(json, 'newPassword') ? undefined : json['newPassword'],
        'autoGeneratePassword': !(0, runtime_1.exists)(json, 'autoGeneratePassword') ? undefined : json['autoGeneratePassword'],
        'loginName': !(0, runtime_1.exists)(json, 'loginName') ? undefined : json['loginName'],
        'domainCode': !(0, runtime_1.exists)(json, 'domainCode') ? undefined : json['domainCode'],
        'locked': !(0, runtime_1.exists)(json, 'locked') ? undefined : json['locked'],
        'inactive': !(0, runtime_1.exists)(json, 'inactive') ? undefined : json['inactive'],
    };
}
exports.WebUserAccountTypeFromJSONTyped = WebUserAccountTypeFromJSONTyped;
function WebUserAccountTypeToJSON(value) {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        'securityQuestion': (0, CodeDescriptionType_1.CodeDescriptionTypeToJSON)(value.securityQuestion),
        'securityAnswer': value.securityAnswer,
        'comments': value.comments,
        'history': (0, WebUserHistoryType_1.WebUserHistoryTypeToJSON)(value.history),
        'newLoginName': value.newLoginName,
        'newPassword': value.newPassword,
        'autoGeneratePassword': value.autoGeneratePassword,
        'loginName': value.loginName,
        'domainCode': value.domainCode,
        'locked': value.locked,
        'inactive': value.inactive,
    };
}
exports.WebUserAccountTypeToJSON = WebUserAccountTypeToJSON;
