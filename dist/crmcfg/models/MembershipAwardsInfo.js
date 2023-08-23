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
exports.MembershipAwardsInfoToJSON = exports.MembershipAwardsInfoFromJSONTyped = exports.MembershipAwardsInfoFromJSON = exports.instanceOfMembershipAwardsInfo = void 0;
const runtime_1 = require("../runtime");
const InstanceLink_1 = require("./InstanceLink");
const MembershipAwardType_1 = require("./MembershipAwardType");
const WarningType_1 = require("./WarningType");
/**
 * Check if a given object implements the MembershipAwardsInfo interface.
 */
function instanceOfMembershipAwardsInfo(value) {
    let isInstance = true;
    return isInstance;
}
exports.instanceOfMembershipAwardsInfo = instanceOfMembershipAwardsInfo;
function MembershipAwardsInfoFromJSON(json) {
    return MembershipAwardsInfoFromJSONTyped(json, false);
}
exports.MembershipAwardsInfoFromJSON = MembershipAwardsInfoFromJSON;
function MembershipAwardsInfoFromJSONTyped(json, ignoreDiscriminator) {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        'membershipAwards': !(0, runtime_1.exists)(json, 'membershipAwards') ? undefined : (json['membershipAwards'].map(MembershipAwardType_1.MembershipAwardTypeFromJSON)),
        'links': !(0, runtime_1.exists)(json, 'links') ? undefined : (json['links'].map(InstanceLink_1.InstanceLinkFromJSON)),
        'warnings': !(0, runtime_1.exists)(json, 'warnings') ? undefined : (json['warnings'].map(WarningType_1.WarningTypeFromJSON)),
    };
}
exports.MembershipAwardsInfoFromJSONTyped = MembershipAwardsInfoFromJSONTyped;
function MembershipAwardsInfoToJSON(value) {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        'membershipAwards': value.membershipAwards === undefined ? undefined : (value.membershipAwards.map(MembershipAwardType_1.MembershipAwardTypeToJSON)),
        'links': value.links === undefined ? undefined : (value.links.map(InstanceLink_1.InstanceLinkToJSON)),
        'warnings': value.warnings === undefined ? undefined : (value.warnings.map(WarningType_1.WarningTypeToJSON)),
    };
}
exports.MembershipAwardsInfoToJSON = MembershipAwardsInfoToJSON;
