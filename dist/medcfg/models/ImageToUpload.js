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
exports.ImageToUploadToJSON = exports.ImageToUploadFromJSONTyped = exports.ImageToUploadFromJSON = exports.instanceOfImageToUpload = void 0;
const runtime_1 = require("../runtime");
const ImageUploadType_1 = require("./ImageUploadType");
const InstanceLink_1 = require("./InstanceLink");
const WarningType_1 = require("./WarningType");
/**
 * Check if a given object implements the ImageToUpload interface.
 */
function instanceOfImageToUpload(value) {
    let isInstance = true;
    return isInstance;
}
exports.instanceOfImageToUpload = instanceOfImageToUpload;
function ImageToUploadFromJSON(json) {
    return ImageToUploadFromJSONTyped(json, false);
}
exports.ImageToUploadFromJSON = ImageToUploadFromJSON;
function ImageToUploadFromJSONTyped(json, ignoreDiscriminator) {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        'images': !(0, runtime_1.exists)(json, 'images') ? undefined : (json['images'].map(ImageUploadType_1.ImageUploadTypeFromJSON)),
        'links': !(0, runtime_1.exists)(json, 'links') ? undefined : (json['links'].map(InstanceLink_1.InstanceLinkFromJSON)),
        'warnings': !(0, runtime_1.exists)(json, 'warnings') ? undefined : (json['warnings'].map(WarningType_1.WarningTypeFromJSON)),
    };
}
exports.ImageToUploadFromJSONTyped = ImageToUploadFromJSONTyped;
function ImageToUploadToJSON(value) {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        'images': value.images === undefined ? undefined : (value.images.map(ImageUploadType_1.ImageUploadTypeToJSON)),
        'links': value.links === undefined ? undefined : (value.links.map(InstanceLink_1.InstanceLinkToJSON)),
        'warnings': value.warnings === undefined ? undefined : (value.warnings.map(WarningType_1.WarningTypeToJSON)),
    };
}
exports.ImageToUploadToJSON = ImageToUploadToJSON;
