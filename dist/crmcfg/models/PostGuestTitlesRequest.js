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
exports.PostGuestTitlesRequestToJSON = exports.PostGuestTitlesRequestFromJSONTyped = exports.PostGuestTitlesRequestFromJSON = exports.instanceOfPostGuestTitlesRequest = void 0;
const runtime_1 = require("../runtime");
const GuestTitleType_1 = require("./GuestTitleType");
const InstanceLink_1 = require("./InstanceLink");
const WarningType_1 = require("./WarningType");
/**
 * Check if a given object implements the PostGuestTitlesRequest interface.
 */
function instanceOfPostGuestTitlesRequest(value) {
    let isInstance = true;
    return isInstance;
}
exports.instanceOfPostGuestTitlesRequest = instanceOfPostGuestTitlesRequest;
function PostGuestTitlesRequestFromJSON(json) {
    return PostGuestTitlesRequestFromJSONTyped(json, false);
}
exports.PostGuestTitlesRequestFromJSON = PostGuestTitlesRequestFromJSON;
function PostGuestTitlesRequestFromJSONTyped(json, ignoreDiscriminator) {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        'guestTitles': !(0, runtime_1.exists)(json, 'guestTitles') ? undefined : (json['guestTitles'].map(GuestTitleType_1.GuestTitleTypeFromJSON)),
        'links': !(0, runtime_1.exists)(json, 'links') ? undefined : (json['links'].map(InstanceLink_1.InstanceLinkFromJSON)),
        'warnings': !(0, runtime_1.exists)(json, 'warnings') ? undefined : (json['warnings'].map(WarningType_1.WarningTypeFromJSON)),
    };
}
exports.PostGuestTitlesRequestFromJSONTyped = PostGuestTitlesRequestFromJSONTyped;
function PostGuestTitlesRequestToJSON(value) {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        'guestTitles': value.guestTitles === undefined ? undefined : (value.guestTitles.map(GuestTitleType_1.GuestTitleTypeToJSON)),
        'links': value.links === undefined ? undefined : (value.links.map(InstanceLink_1.InstanceLinkToJSON)),
        'warnings': value.warnings === undefined ? undefined : (value.warnings.map(WarningType_1.WarningTypeToJSON)),
    };
}
exports.PostGuestTitlesRequestToJSON = PostGuestTitlesRequestToJSON;
