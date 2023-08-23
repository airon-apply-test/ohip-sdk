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
exports.OwnerTeamsCriteriaToJSON = exports.OwnerTeamsCriteriaFromJSONTyped = exports.OwnerTeamsCriteriaFromJSON = exports.instanceOfOwnerTeamsCriteria = void 0;
const runtime_1 = require("../runtime");
const InstanceLink_1 = require("./InstanceLink");
const OwnerTeamType_1 = require("./OwnerTeamType");
const WarningType_1 = require("./WarningType");
/**
 * Check if a given object implements the OwnerTeamsCriteria interface.
 */
function instanceOfOwnerTeamsCriteria(value) {
    let isInstance = true;
    return isInstance;
}
exports.instanceOfOwnerTeamsCriteria = instanceOfOwnerTeamsCriteria;
function OwnerTeamsCriteriaFromJSON(json) {
    return OwnerTeamsCriteriaFromJSONTyped(json, false);
}
exports.OwnerTeamsCriteriaFromJSON = OwnerTeamsCriteriaFromJSON;
function OwnerTeamsCriteriaFromJSONTyped(json, ignoreDiscriminator) {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        'ownerTeams': !(0, runtime_1.exists)(json, 'ownerTeams') ? undefined : (json['ownerTeams'].map(OwnerTeamType_1.OwnerTeamTypeFromJSON)),
        'links': !(0, runtime_1.exists)(json, 'links') ? undefined : (json['links'].map(InstanceLink_1.InstanceLinkFromJSON)),
        'warnings': !(0, runtime_1.exists)(json, 'warnings') ? undefined : (json['warnings'].map(WarningType_1.WarningTypeFromJSON)),
    };
}
exports.OwnerTeamsCriteriaFromJSONTyped = OwnerTeamsCriteriaFromJSONTyped;
function OwnerTeamsCriteriaToJSON(value) {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        'ownerTeams': value.ownerTeams === undefined ? undefined : (value.ownerTeams.map(OwnerTeamType_1.OwnerTeamTypeToJSON)),
        'links': value.links === undefined ? undefined : (value.links.map(InstanceLink_1.InstanceLinkToJSON)),
        'warnings': value.warnings === undefined ? undefined : (value.warnings.map(WarningType_1.WarningTypeToJSON)),
    };
}
exports.OwnerTeamsCriteriaToJSON = OwnerTeamsCriteriaToJSON;
