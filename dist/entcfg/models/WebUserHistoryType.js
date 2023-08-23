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
exports.WebUserHistoryTypeToJSON = exports.WebUserHistoryTypeFromJSONTyped = exports.WebUserHistoryTypeFromJSON = exports.instanceOfWebUserHistoryType = void 0;
const runtime_1 = require("../runtime");
/**
 * Check if a given object implements the WebUserHistoryType interface.
 */
function instanceOfWebUserHistoryType(value) {
    let isInstance = true;
    return isInstance;
}
exports.instanceOfWebUserHistoryType = instanceOfWebUserHistoryType;
function WebUserHistoryTypeFromJSON(json) {
    return WebUserHistoryTypeFromJSONTyped(json, false);
}
exports.WebUserHistoryTypeFromJSON = WebUserHistoryTypeFromJSON;
function WebUserHistoryTypeFromJSONTyped(json, ignoreDiscriminator) {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        'createDateTime': !(0, runtime_1.exists)(json, 'createDateTime') ? undefined : (new Date(json['createDateTime'])),
        'creatorId': !(0, runtime_1.exists)(json, 'creatorId') ? undefined : json['creatorId'],
        'lastModifyDateTime': !(0, runtime_1.exists)(json, 'lastModifyDateTime') ? undefined : (new Date(json['lastModifyDateTime'])),
        'lastModifierId': !(0, runtime_1.exists)(json, 'lastModifierId') ? undefined : json['lastModifierId'],
        'purgeDate': !(0, runtime_1.exists)(json, 'purgeDate') ? undefined : (new Date(json['purgeDate'])),
        'lastLogin': !(0, runtime_1.exists)(json, 'lastLogin') ? undefined : (new Date(json['lastLogin'])),
        'lastPasswordChange': !(0, runtime_1.exists)(json, 'lastPasswordChange') ? undefined : (new Date(json['lastPasswordChange'])),
        'inactiveDate': !(0, runtime_1.exists)(json, 'inactiveDate') ? undefined : (new Date(json['inactiveDate'])),
    };
}
exports.WebUserHistoryTypeFromJSONTyped = WebUserHistoryTypeFromJSONTyped;
function WebUserHistoryTypeToJSON(value) {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        'createDateTime': value.createDateTime === undefined ? undefined : (value.createDateTime.toISOString()),
        'creatorId': value.creatorId,
        'lastModifyDateTime': value.lastModifyDateTime === undefined ? undefined : (value.lastModifyDateTime.toISOString()),
        'lastModifierId': value.lastModifierId,
        'purgeDate': value.purgeDate === undefined ? undefined : (value.purgeDate.toISOString().substr(0, 10)),
        'lastLogin': value.lastLogin === undefined ? undefined : (value.lastLogin.toISOString().substr(0, 10)),
        'lastPasswordChange': value.lastPasswordChange === undefined ? undefined : (value.lastPasswordChange.toISOString().substr(0, 10)),
        'inactiveDate': value.inactiveDate === undefined ? undefined : (value.inactiveDate.toISOString().substr(0, 10)),
    };
}
exports.WebUserHistoryTypeToJSON = WebUserHistoryTypeToJSON;
