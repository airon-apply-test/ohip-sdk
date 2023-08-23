"use strict";
/* tslint:disable */
/* eslint-disable */
/**
 * OPERA Cloud Content Service
 * Opera Content Service offers capability to manage large content such as images and files.<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 22.3.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
Object.defineProperty(exports, "__esModule", { value: true });
exports.EmailFileToJSON = exports.EmailFileFromJSONTyped = exports.EmailFileFromJSON = exports.instanceOfEmailFile = void 0;
const runtime_1 = require("../runtime");
const EmailFileType_1 = require("./EmailFileType");
const InstanceLink_1 = require("./InstanceLink");
const WarningType_1 = require("./WarningType");
/**
 * Check if a given object implements the EmailFile interface.
 */
function instanceOfEmailFile(value) {
    let isInstance = true;
    return isInstance;
}
exports.instanceOfEmailFile = instanceOfEmailFile;
function EmailFileFromJSON(json) {
    return EmailFileFromJSONTyped(json, false);
}
exports.EmailFileFromJSON = EmailFileFromJSON;
function EmailFileFromJSONTyped(json, ignoreDiscriminator) {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        'file': !(0, runtime_1.exists)(json, 'file') ? undefined : (0, EmailFileType_1.EmailFileTypeFromJSON)(json['file']),
        'links': !(0, runtime_1.exists)(json, 'links') ? undefined : (json['links'].map(InstanceLink_1.InstanceLinkFromJSON)),
        'warnings': !(0, runtime_1.exists)(json, 'warnings') ? undefined : (json['warnings'].map(WarningType_1.WarningTypeFromJSON)),
    };
}
exports.EmailFileFromJSONTyped = EmailFileFromJSONTyped;
function EmailFileToJSON(value) {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        'file': (0, EmailFileType_1.EmailFileTypeToJSON)(value.file),
        'links': value.links === undefined ? undefined : (value.links.map(InstanceLink_1.InstanceLinkToJSON)),
        'warnings': value.warnings === undefined ? undefined : (value.warnings.map(WarningType_1.WarningTypeToJSON)),
    };
}
exports.EmailFileToJSON = EmailFileToJSON;
