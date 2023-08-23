"use strict";
/* tslint:disable */
/* eslint-disable */
/**
 * OPERA Cloud Accounts Receivables API
 * APIs to cater for Accounts Receivables functionality in OPERA Cloud. <br /><br The OPERA Cloud Accounts Receivable module enables you to manage debtors’ accounts, invoices, and remittance.<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 22.3.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
Object.defineProperty(exports, "__esModule", { value: true });
exports.CommentTypeToJSON = exports.CommentTypeFromJSONTyped = exports.CommentTypeFromJSON = exports.instanceOfCommentType = void 0;
const runtime_1 = require("../runtime");
const FormattedTextTextType_1 = require("./FormattedTextTextType");
/**
 * Check if a given object implements the CommentType interface.
 */
function instanceOfCommentType(value) {
    let isInstance = true;
    return isInstance;
}
exports.instanceOfCommentType = instanceOfCommentType;
function CommentTypeFromJSON(json) {
    return CommentTypeFromJSONTyped(json, false);
}
exports.CommentTypeFromJSON = CommentTypeFromJSON;
function CommentTypeFromJSONTyped(json, ignoreDiscriminator) {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        'text': !(0, runtime_1.exists)(json, 'text') ? undefined : (0, FormattedTextTextType_1.FormattedTextTextTypeFromJSON)(json['text']),
        'image': !(0, runtime_1.exists)(json, 'image') ? undefined : json['image'],
        'url': !(0, runtime_1.exists)(json, 'url') ? undefined : json['url'],
        'commentTitle': !(0, runtime_1.exists)(json, 'commentTitle') ? undefined : json['commentTitle'],
        'notificationLocation': !(0, runtime_1.exists)(json, 'notificationLocation') ? undefined : json['notificationLocation'],
        'type': !(0, runtime_1.exists)(json, 'type') ? undefined : json['type'],
        'typeDescription': !(0, runtime_1.exists)(json, 'typeDescription') ? undefined : json['typeDescription'],
        'internal': !(0, runtime_1.exists)(json, 'internal') ? undefined : json['internal'],
        'confidential': !(0, runtime_1.exists)(json, 'confidential') ? undefined : json['confidential'],
        'overrideInternal': !(0, runtime_1.exists)(json, 'overrideInternal') ? undefined : json['overrideInternal'],
        'protectDescription': !(0, runtime_1.exists)(json, 'protectDescription') ? undefined : json['protectDescription'],
        'hotelId': !(0, runtime_1.exists)(json, 'hotelId') ? undefined : json['hotelId'],
        'actionType': !(0, runtime_1.exists)(json, 'actionType') ? undefined : json['actionType'],
        'actionDate': !(0, runtime_1.exists)(json, 'actionDate') ? undefined : (new Date(json['actionDate'])),
        'createDateTime': !(0, runtime_1.exists)(json, 'createDateTime') ? undefined : (new Date(json['createDateTime'])),
        'creatorId': !(0, runtime_1.exists)(json, 'creatorId') ? undefined : json['creatorId'],
        'lastModifyDateTime': !(0, runtime_1.exists)(json, 'lastModifyDateTime') ? undefined : (new Date(json['lastModifyDateTime'])),
        'lastModifierId': !(0, runtime_1.exists)(json, 'lastModifierId') ? undefined : json['lastModifierId'],
        'purgeDate': !(0, runtime_1.exists)(json, 'purgeDate') ? undefined : (new Date(json['purgeDate'])),
    };
}
exports.CommentTypeFromJSONTyped = CommentTypeFromJSONTyped;
function CommentTypeToJSON(value) {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        'text': (0, FormattedTextTextType_1.FormattedTextTextTypeToJSON)(value.text),
        'image': value.image,
        'url': value.url,
        'commentTitle': value.commentTitle,
        'notificationLocation': value.notificationLocation,
        'type': value.type,
        'typeDescription': value.typeDescription,
        'internal': value.internal,
        'confidential': value.confidential,
        'overrideInternal': value.overrideInternal,
        'protectDescription': value.protectDescription,
        'hotelId': value.hotelId,
        'actionType': value.actionType,
        'actionDate': value.actionDate === undefined ? undefined : (value.actionDate.toISOString().substr(0, 10)),
        'createDateTime': value.createDateTime === undefined ? undefined : (value.createDateTime.toISOString()),
        'creatorId': value.creatorId,
        'lastModifyDateTime': value.lastModifyDateTime === undefined ? undefined : (value.lastModifyDateTime.toISOString()),
        'lastModifierId': value.lastModifierId,
        'purgeDate': value.purgeDate === undefined ? undefined : (value.purgeDate.toISOString().substr(0, 10)),
    };
}
exports.CommentTypeToJSON = CommentTypeToJSON;
