"use strict";
/* tslint:disable */
/* eslint-disable */
/**
 * OPERA Cloud Price Availability Rate API
 * APIs to cater for Price and Rate Availability functionality in OPERA Cloud. <br /><br />Availability enables you to manage your room inventory by providing a detailed view of all available and sold rooms at a property. Some of the tasks you can perform include defining conditions for stay restrictions, setting room sell limits, and searching for and viewing room availability.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 21.5.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
Object.defineProperty(exports, "__esModule", { value: true });
exports.EmailTypeToJSON = exports.EmailTypeFromJSONTyped = exports.EmailTypeFromJSON = exports.instanceOfEmailType = exports.EmailTypeEmailFormatEnum = void 0;
const runtime_1 = require("../runtime");
/**
 * @export
 */
exports.EmailTypeEmailFormatEnum = {
    Html: 'Html',
    Text: 'Text'
};
/**
 * Check if a given object implements the EmailType interface.
 */
function instanceOfEmailType(value) {
    let isInstance = true;
    return isInstance;
}
exports.instanceOfEmailType = instanceOfEmailType;
function EmailTypeFromJSON(json) {
    return EmailTypeFromJSONTyped(json, false);
}
exports.EmailTypeFromJSON = EmailTypeFromJSON;
function EmailTypeFromJSONTyped(json, ignoreDiscriminator) {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        'emailAddress': !(0, runtime_1.exists)(json, 'emailAddress') ? undefined : json['emailAddress'],
        'type': !(0, runtime_1.exists)(json, 'type') ? undefined : json['type'],
        'typeDescription': !(0, runtime_1.exists)(json, 'typeDescription') ? undefined : json['typeDescription'],
        'emailFormat': !(0, runtime_1.exists)(json, 'emailFormat') ? undefined : json['emailFormat'],
        'primaryInd': !(0, runtime_1.exists)(json, 'primaryInd') ? undefined : json['primaryInd'],
        'orderSequence': !(0, runtime_1.exists)(json, 'orderSequence') ? undefined : json['orderSequence'],
        'createDateTime': !(0, runtime_1.exists)(json, 'createDateTime') ? undefined : (new Date(json['createDateTime'])),
        'creatorId': !(0, runtime_1.exists)(json, 'creatorId') ? undefined : json['creatorId'],
        'lastModifyDateTime': !(0, runtime_1.exists)(json, 'lastModifyDateTime') ? undefined : (new Date(json['lastModifyDateTime'])),
        'lastModifierId': !(0, runtime_1.exists)(json, 'lastModifierId') ? undefined : json['lastModifierId'],
        'purgeDate': !(0, runtime_1.exists)(json, 'purgeDate') ? undefined : (new Date(json['purgeDate'])),
    };
}
exports.EmailTypeFromJSONTyped = EmailTypeFromJSONTyped;
function EmailTypeToJSON(value) {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        'emailAddress': value.emailAddress,
        'type': value.type,
        'typeDescription': value.typeDescription,
        'emailFormat': value.emailFormat,
        'primaryInd': value.primaryInd,
        'orderSequence': value.orderSequence,
        'createDateTime': value.createDateTime === undefined ? undefined : (value.createDateTime.toISOString()),
        'creatorId': value.creatorId,
        'lastModifyDateTime': value.lastModifyDateTime === undefined ? undefined : (value.lastModifyDateTime.toISOString()),
        'lastModifierId': value.lastModifierId,
        'purgeDate': value.purgeDate === undefined ? undefined : (value.purgeDate.toISOString().substr(0, 10)),
    };
}
exports.EmailTypeToJSON = EmailTypeToJSON;
