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
exports.MembershipTypeRulePointsDefinitionTypeToJSON = exports.MembershipTypeRulePointsDefinitionTypeFromJSONTyped = exports.MembershipTypeRulePointsDefinitionTypeFromJSON = exports.instanceOfMembershipTypeRulePointsDefinitionType = void 0;
const runtime_1 = require("../runtime");
const PointsRoundingFlagType_1 = require("./PointsRoundingFlagType");
/**
 * Check if a given object implements the MembershipTypeRulePointsDefinitionType interface.
 */
function instanceOfMembershipTypeRulePointsDefinitionType(value) {
    let isInstance = true;
    return isInstance;
}
exports.instanceOfMembershipTypeRulePointsDefinitionType = instanceOfMembershipTypeRulePointsDefinitionType;
function MembershipTypeRulePointsDefinitionTypeFromJSON(json) {
    return MembershipTypeRulePointsDefinitionTypeFromJSONTyped(json, false);
}
exports.MembershipTypeRulePointsDefinitionTypeFromJSON = MembershipTypeRulePointsDefinitionTypeFromJSON;
function MembershipTypeRulePointsDefinitionTypeFromJSONTyped(json, ignoreDiscriminator) {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        'points': !(0, runtime_1.exists)(json, 'points') ? undefined : json['points'],
        'costPerPoint': !(0, runtime_1.exists)(json, 'costPerPoint') ? undefined : json['costPerPoint'],
        'revenueUnits': !(0, runtime_1.exists)(json, 'revenueUnits') ? undefined : json['revenueUnits'],
        'pointsRoundingFlag': !(0, runtime_1.exists)(json, 'pointsRoundingFlag') ? undefined : (0, PointsRoundingFlagType_1.PointsRoundingFlagTypeFromJSON)(json['pointsRoundingFlag']),
        'percentagePoints': !(0, runtime_1.exists)(json, 'percentagePoints') ? undefined : json['percentagePoints'],
        'excludePointProjection': !(0, runtime_1.exists)(json, 'excludePointProjection') ? undefined : json['excludePointProjection'],
    };
}
exports.MembershipTypeRulePointsDefinitionTypeFromJSONTyped = MembershipTypeRulePointsDefinitionTypeFromJSONTyped;
function MembershipTypeRulePointsDefinitionTypeToJSON(value) {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        'points': value.points,
        'costPerPoint': value.costPerPoint,
        'revenueUnits': value.revenueUnits,
        'pointsRoundingFlag': (0, PointsRoundingFlagType_1.PointsRoundingFlagTypeToJSON)(value.pointsRoundingFlag),
        'percentagePoints': value.percentagePoints,
        'excludePointProjection': value.excludePointProjection,
    };
}
exports.MembershipTypeRulePointsDefinitionTypeToJSON = MembershipTypeRulePointsDefinitionTypeToJSON;
