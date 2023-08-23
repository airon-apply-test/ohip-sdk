"use strict";
/* tslint:disable */
/* eslint-disable */
/**
 * OPERA Cloud CRM Configuration API
 * APIs for Customer Relationship Management (profile) configuration, such as creating preferences, or address types.  It also includes Membership Configuration, where you can retrieve membership levels that are configured for a property, or create new membership enrollment codes.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 21.5.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
Object.defineProperty(exports, "__esModule", { value: true });
exports.PostGuestTypesRequestToJSON = exports.PostGuestTypesRequestFromJSONTyped = exports.PostGuestTypesRequestFromJSON = exports.instanceOfPostGuestTypesRequest = void 0;
const runtime_1 = require("../runtime");
const GuestTypeType_1 = require("./GuestTypeType");
const InstanceLink_1 = require("./InstanceLink");
const WarningType_1 = require("./WarningType");
/**
 * Check if a given object implements the PostGuestTypesRequest interface.
 */
function instanceOfPostGuestTypesRequest(value) {
    let isInstance = true;
    return isInstance;
}
exports.instanceOfPostGuestTypesRequest = instanceOfPostGuestTypesRequest;
function PostGuestTypesRequestFromJSON(json) {
    return PostGuestTypesRequestFromJSONTyped(json, false);
}
exports.PostGuestTypesRequestFromJSON = PostGuestTypesRequestFromJSON;
function PostGuestTypesRequestFromJSONTyped(json, ignoreDiscriminator) {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        'guestTypes': !(0, runtime_1.exists)(json, 'guestTypes') ? undefined : (json['guestTypes'].map(GuestTypeType_1.GuestTypeTypeFromJSON)),
        'links': !(0, runtime_1.exists)(json, 'links') ? undefined : (json['links'].map(InstanceLink_1.InstanceLinkFromJSON)),
        'warnings': !(0, runtime_1.exists)(json, 'warnings') ? undefined : (json['warnings'].map(WarningType_1.WarningTypeFromJSON)),
    };
}
exports.PostGuestTypesRequestFromJSONTyped = PostGuestTypesRequestFromJSONTyped;
function PostGuestTypesRequestToJSON(value) {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        'guestTypes': value.guestTypes === undefined ? undefined : (value.guestTypes.map(GuestTypeType_1.GuestTypeTypeToJSON)),
        'links': value.links === undefined ? undefined : (value.links.map(InstanceLink_1.InstanceLinkToJSON)),
        'warnings': value.warnings === undefined ? undefined : (value.warnings.map(WarningType_1.WarningTypeToJSON)),
    };
}
exports.PostGuestTypesRequestToJSON = PostGuestTypesRequestToJSON;
