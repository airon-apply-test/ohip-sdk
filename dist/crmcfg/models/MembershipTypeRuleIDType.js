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
exports.MembershipTypeRuleIDTypeToJSON = exports.MembershipTypeRuleIDTypeFromJSONTyped = exports.MembershipTypeRuleIDTypeFromJSON = exports.instanceOfMembershipTypeRuleIDType = void 0;
const runtime_1 = require("../runtime");
/**
 * Check if a given object implements the MembershipTypeRuleIDType interface.
 */
function instanceOfMembershipTypeRuleIDType(value) {
    let isInstance = true;
    return isInstance;
}
exports.instanceOfMembershipTypeRuleIDType = instanceOfMembershipTypeRuleIDType;
function MembershipTypeRuleIDTypeFromJSON(json) {
    return MembershipTypeRuleIDTypeFromJSONTyped(json, false);
}
exports.MembershipTypeRuleIDTypeFromJSON = MembershipTypeRuleIDTypeFromJSON;
function MembershipTypeRuleIDTypeFromJSONTyped(json, ignoreDiscriminator) {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        'membershipType': !(0, runtime_1.exists)(json, 'membershipType') ? undefined : json['membershipType'],
        'membershipPointsSequence': !(0, runtime_1.exists)(json, 'membershipPointsSequence') ? undefined : json['membershipPointsSequence'],
    };
}
exports.MembershipTypeRuleIDTypeFromJSONTyped = MembershipTypeRuleIDTypeFromJSONTyped;
function MembershipTypeRuleIDTypeToJSON(value) {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        'membershipType': value.membershipType,
        'membershipPointsSequence': value.membershipPointsSequence,
    };
}
exports.MembershipTypeRuleIDTypeToJSON = MembershipTypeRuleIDTypeToJSON;
