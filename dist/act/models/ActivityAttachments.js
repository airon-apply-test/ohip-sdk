"use strict";
/* tslint:disable */
/* eslint-disable */
/**
 * OPERA Cloud Activity API
 * APIs to cater for Sales Activity functionality in OPERA Cloud. <br /><br /> Activities provide you with an account management tool for managing daily tasks such as appointments, sales calls, contact follow-up, and so on.<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 22.3.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
Object.defineProperty(exports, "__esModule", { value: true });
exports.ActivityAttachmentsToJSON = exports.ActivityAttachmentsFromJSONTyped = exports.ActivityAttachmentsFromJSON = exports.instanceOfActivityAttachments = void 0;
const runtime_1 = require("../runtime");
const AttachmentType_1 = require("./AttachmentType");
const InstanceLink_1 = require("./InstanceLink");
const WarningType_1 = require("./WarningType");
/**
 * Check if a given object implements the ActivityAttachments interface.
 */
function instanceOfActivityAttachments(value) {
    let isInstance = true;
    return isInstance;
}
exports.instanceOfActivityAttachments = instanceOfActivityAttachments;
function ActivityAttachmentsFromJSON(json) {
    return ActivityAttachmentsFromJSONTyped(json, false);
}
exports.ActivityAttachmentsFromJSON = ActivityAttachmentsFromJSON;
function ActivityAttachmentsFromJSONTyped(json, ignoreDiscriminator) {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        'activityAttachmentList': !(0, runtime_1.exists)(json, 'activityAttachmentList') ? undefined : (json['activityAttachmentList'].map(AttachmentType_1.AttachmentTypeFromJSON)),
        'links': !(0, runtime_1.exists)(json, 'links') ? undefined : (json['links'].map(InstanceLink_1.InstanceLinkFromJSON)),
        'warnings': !(0, runtime_1.exists)(json, 'warnings') ? undefined : (json['warnings'].map(WarningType_1.WarningTypeFromJSON)),
    };
}
exports.ActivityAttachmentsFromJSONTyped = ActivityAttachmentsFromJSONTyped;
function ActivityAttachmentsToJSON(value) {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        'activityAttachmentList': value.activityAttachmentList === undefined ? undefined : (value.activityAttachmentList.map(AttachmentType_1.AttachmentTypeToJSON)),
        'links': value.links === undefined ? undefined : (value.links.map(InstanceLink_1.InstanceLinkToJSON)),
        'warnings': value.warnings === undefined ? undefined : (value.warnings.map(WarningType_1.WarningTypeToJSON)),
    };
}
exports.ActivityAttachmentsToJSON = ActivityAttachmentsToJSON;
