"use strict";
/* tslint:disable */
/* eslint-disable */
/**
 * OPERA Cloud API for Customer Management Service
 * This API deals with the different aspect of the CustomerManagement.<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 22.3.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
Object.defineProperty(exports, "__esModule", { value: true });
exports.CalendarTaskAttachmentsToJSON = exports.CalendarTaskAttachmentsFromJSONTyped = exports.CalendarTaskAttachmentsFromJSON = exports.instanceOfCalendarTaskAttachments = void 0;
const runtime_1 = require("../runtime");
const AttachmentType_1 = require("./AttachmentType");
const InstanceLink_1 = require("./InstanceLink");
const WarningType_1 = require("./WarningType");
/**
 * Check if a given object implements the CalendarTaskAttachments interface.
 */
function instanceOfCalendarTaskAttachments(value) {
    let isInstance = true;
    return isInstance;
}
exports.instanceOfCalendarTaskAttachments = instanceOfCalendarTaskAttachments;
function CalendarTaskAttachmentsFromJSON(json) {
    return CalendarTaskAttachmentsFromJSONTyped(json, false);
}
exports.CalendarTaskAttachmentsFromJSON = CalendarTaskAttachmentsFromJSON;
function CalendarTaskAttachmentsFromJSONTyped(json, ignoreDiscriminator) {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        'calendarTaskAttachments': !(0, runtime_1.exists)(json, 'calendarTaskAttachments') ? undefined : (json['calendarTaskAttachments'].map(AttachmentType_1.AttachmentTypeFromJSON)),
        'links': !(0, runtime_1.exists)(json, 'links') ? undefined : (json['links'].map(InstanceLink_1.InstanceLinkFromJSON)),
        'warnings': !(0, runtime_1.exists)(json, 'warnings') ? undefined : (json['warnings'].map(WarningType_1.WarningTypeFromJSON)),
    };
}
exports.CalendarTaskAttachmentsFromJSONTyped = CalendarTaskAttachmentsFromJSONTyped;
function CalendarTaskAttachmentsToJSON(value) {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        'calendarTaskAttachments': value.calendarTaskAttachments === undefined ? undefined : (value.calendarTaskAttachments.map(AttachmentType_1.AttachmentTypeToJSON)),
        'links': value.links === undefined ? undefined : (value.links.map(InstanceLink_1.InstanceLinkToJSON)),
        'warnings': value.warnings === undefined ? undefined : (value.warnings.map(WarningType_1.WarningTypeToJSON)),
    };
}
exports.CalendarTaskAttachmentsToJSON = CalendarTaskAttachmentsToJSON;
