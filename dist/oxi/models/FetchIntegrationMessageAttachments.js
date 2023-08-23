"use strict";
/* tslint:disable */
/* eslint-disable */
/**
 * OPERA Cloud Xchange Interface OXI API
 * APIs to create and manage OPERA Xchange Interface (OXI) configurations for OPERA Cloud Exchange module functionality.<br /><br /> Compatible with OPERA Cloud release 23.0.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 23.0.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
Object.defineProperty(exports, "__esModule", { value: true });
exports.FetchIntegrationMessageAttachmentsToJSON = exports.FetchIntegrationMessageAttachmentsFromJSONTyped = exports.FetchIntegrationMessageAttachmentsFromJSON = exports.instanceOfFetchIntegrationMessageAttachments = void 0;
const runtime_1 = require("../runtime");
const InstanceLink_1 = require("./InstanceLink");
const IntegrationMessageAttachmentsType_1 = require("./IntegrationMessageAttachmentsType");
const WarningType_1 = require("./WarningType");
/**
 * Check if a given object implements the FetchIntegrationMessageAttachments interface.
 */
function instanceOfFetchIntegrationMessageAttachments(value) {
    let isInstance = true;
    return isInstance;
}
exports.instanceOfFetchIntegrationMessageAttachments = instanceOfFetchIntegrationMessageAttachments;
function FetchIntegrationMessageAttachmentsFromJSON(json) {
    return FetchIntegrationMessageAttachmentsFromJSONTyped(json, false);
}
exports.FetchIntegrationMessageAttachmentsFromJSON = FetchIntegrationMessageAttachmentsFromJSON;
function FetchIntegrationMessageAttachmentsFromJSONTyped(json, ignoreDiscriminator) {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        'message': !(0, runtime_1.exists)(json, 'message') ? undefined : (0, IntegrationMessageAttachmentsType_1.IntegrationMessageAttachmentsTypeFromJSON)(json['message']),
        'links': !(0, runtime_1.exists)(json, 'links') ? undefined : (json['links'].map(InstanceLink_1.InstanceLinkFromJSON)),
        'warnings': !(0, runtime_1.exists)(json, 'warnings') ? undefined : (json['warnings'].map(WarningType_1.WarningTypeFromJSON)),
    };
}
exports.FetchIntegrationMessageAttachmentsFromJSONTyped = FetchIntegrationMessageAttachmentsFromJSONTyped;
function FetchIntegrationMessageAttachmentsToJSON(value) {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        'message': (0, IntegrationMessageAttachmentsType_1.IntegrationMessageAttachmentsTypeToJSON)(value.message),
        'links': value.links === undefined ? undefined : (value.links.map(InstanceLink_1.InstanceLinkToJSON)),
        'warnings': value.warnings === undefined ? undefined : (value.warnings.map(WarningType_1.WarningTypeToJSON)),
    };
}
exports.FetchIntegrationMessageAttachmentsToJSON = FetchIntegrationMessageAttachmentsToJSON;
