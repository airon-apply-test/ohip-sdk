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
exports.PostMembershipTypesRequestToJSON = exports.PostMembershipTypesRequestFromJSONTyped = exports.PostMembershipTypesRequestFromJSON = exports.instanceOfPostMembershipTypesRequest = void 0;
const runtime_1 = require("../runtime");
const InstanceLink_1 = require("./InstanceLink");
const MembershipTypeType_1 = require("./MembershipTypeType");
const WarningType_1 = require("./WarningType");
/**
 * Check if a given object implements the PostMembershipTypesRequest interface.
 */
function instanceOfPostMembershipTypesRequest(value) {
    let isInstance = true;
    return isInstance;
}
exports.instanceOfPostMembershipTypesRequest = instanceOfPostMembershipTypesRequest;
function PostMembershipTypesRequestFromJSON(json) {
    return PostMembershipTypesRequestFromJSONTyped(json, false);
}
exports.PostMembershipTypesRequestFromJSON = PostMembershipTypesRequestFromJSON;
function PostMembershipTypesRequestFromJSONTyped(json, ignoreDiscriminator) {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        'membershipTypes': !(0, runtime_1.exists)(json, 'membershipTypes') ? undefined : (json['membershipTypes'].map(MembershipTypeType_1.MembershipTypeTypeFromJSON)),
        'links': !(0, runtime_1.exists)(json, 'links') ? undefined : (json['links'].map(InstanceLink_1.InstanceLinkFromJSON)),
        'warnings': !(0, runtime_1.exists)(json, 'warnings') ? undefined : (json['warnings'].map(WarningType_1.WarningTypeFromJSON)),
    };
}
exports.PostMembershipTypesRequestFromJSONTyped = PostMembershipTypesRequestFromJSONTyped;
function PostMembershipTypesRequestToJSON(value) {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        'membershipTypes': value.membershipTypes === undefined ? undefined : (value.membershipTypes.map(MembershipTypeType_1.MembershipTypeTypeToJSON)),
        'links': value.links === undefined ? undefined : (value.links.map(InstanceLink_1.InstanceLinkToJSON)),
        'warnings': value.warnings === undefined ? undefined : (value.warnings.map(WarningType_1.WarningTypeToJSON)),
    };
}
exports.PostMembershipTypesRequestToJSON = PostMembershipTypesRequestToJSON;
