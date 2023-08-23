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
exports.TitlesCriteriaToJSON = exports.TitlesCriteriaFromJSONTyped = exports.TitlesCriteriaFromJSON = exports.instanceOfTitlesCriteria = void 0;
const runtime_1 = require("../runtime");
const InstanceLink_1 = require("./InstanceLink");
const TitleType_1 = require("./TitleType");
const WarningType_1 = require("./WarningType");
/**
 * Check if a given object implements the TitlesCriteria interface.
 */
function instanceOfTitlesCriteria(value) {
    let isInstance = true;
    return isInstance;
}
exports.instanceOfTitlesCriteria = instanceOfTitlesCriteria;
function TitlesCriteriaFromJSON(json) {
    return TitlesCriteriaFromJSONTyped(json, false);
}
exports.TitlesCriteriaFromJSON = TitlesCriteriaFromJSON;
function TitlesCriteriaFromJSONTyped(json, ignoreDiscriminator) {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        'titles': !(0, runtime_1.exists)(json, 'titles') ? undefined : (json['titles'].map(TitleType_1.TitleTypeFromJSON)),
        'links': !(0, runtime_1.exists)(json, 'links') ? undefined : (json['links'].map(InstanceLink_1.InstanceLinkFromJSON)),
        'warnings': !(0, runtime_1.exists)(json, 'warnings') ? undefined : (json['warnings'].map(WarningType_1.WarningTypeFromJSON)),
    };
}
exports.TitlesCriteriaFromJSONTyped = TitlesCriteriaFromJSONTyped;
function TitlesCriteriaToJSON(value) {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        'titles': value.titles === undefined ? undefined : (value.titles.map(TitleType_1.TitleTypeToJSON)),
        'links': value.links === undefined ? undefined : (value.links.map(InstanceLink_1.InstanceLinkToJSON)),
        'warnings': value.warnings === undefined ? undefined : (value.warnings.map(WarningType_1.WarningTypeToJSON)),
    };
}
exports.TitlesCriteriaToJSON = TitlesCriteriaToJSON;
