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
exports.MembershipEnrollmentGroupTypeToJSON = exports.MembershipEnrollmentGroupTypeFromJSONTyped = exports.MembershipEnrollmentGroupTypeFromJSON = exports.instanceOfMembershipEnrollmentGroupType = void 0;
const runtime_1 = require("../runtime");
const CodeDescriptionType_1 = require("./CodeDescriptionType");
/**
 * Check if a given object implements the MembershipEnrollmentGroupType interface.
 */
function instanceOfMembershipEnrollmentGroupType(value) {
    let isInstance = true;
    return isInstance;
}
exports.instanceOfMembershipEnrollmentGroupType = instanceOfMembershipEnrollmentGroupType;
function MembershipEnrollmentGroupTypeFromJSON(json) {
    return MembershipEnrollmentGroupTypeFromJSONTyped(json, false);
}
exports.MembershipEnrollmentGroupTypeFromJSON = MembershipEnrollmentGroupTypeFromJSON;
function MembershipEnrollmentGroupTypeFromJSONTyped(json, ignoreDiscriminator) {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        'code': !(0, runtime_1.exists)(json, 'code') ? undefined : json['code'],
        'description': !(0, runtime_1.exists)(json, 'description') ? undefined : json['description'],
        'displaySequence': !(0, runtime_1.exists)(json, 'displaySequence') ? undefined : json['displaySequence'],
        'enrollmentCodes': !(0, runtime_1.exists)(json, 'enrollmentCodes') ? undefined : (json['enrollmentCodes'].map(CodeDescriptionType_1.CodeDescriptionTypeFromJSON)),
    };
}
exports.MembershipEnrollmentGroupTypeFromJSONTyped = MembershipEnrollmentGroupTypeFromJSONTyped;
function MembershipEnrollmentGroupTypeToJSON(value) {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        'code': value.code,
        'description': value.description,
        'displaySequence': value.displaySequence,
        'enrollmentCodes': value.enrollmentCodes === undefined ? undefined : (value.enrollmentCodes.map(CodeDescriptionType_1.CodeDescriptionTypeToJSON)),
    };
}
exports.MembershipEnrollmentGroupTypeToJSON = MembershipEnrollmentGroupTypeToJSON;
