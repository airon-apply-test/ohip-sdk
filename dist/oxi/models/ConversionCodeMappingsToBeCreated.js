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
exports.ConversionCodeMappingsToBeCreatedToJSON = exports.ConversionCodeMappingsToBeCreatedFromJSONTyped = exports.ConversionCodeMappingsToBeCreatedFromJSON = exports.instanceOfConversionCodeMappingsToBeCreated = void 0;
const runtime_1 = require("../runtime");
const ConversionCodeMappingType_1 = require("./ConversionCodeMappingType");
const InstanceLink_1 = require("./InstanceLink");
const IntegrationSystemType_1 = require("./IntegrationSystemType");
const WarningType_1 = require("./WarningType");
/**
 * Check if a given object implements the ConversionCodeMappingsToBeCreated interface.
 */
function instanceOfConversionCodeMappingsToBeCreated(value) {
    let isInstance = true;
    return isInstance;
}
exports.instanceOfConversionCodeMappingsToBeCreated = instanceOfConversionCodeMappingsToBeCreated;
function ConversionCodeMappingsToBeCreatedFromJSON(json) {
    return ConversionCodeMappingsToBeCreatedFromJSONTyped(json, false);
}
exports.ConversionCodeMappingsToBeCreatedFromJSON = ConversionCodeMappingsToBeCreatedFromJSON;
function ConversionCodeMappingsToBeCreatedFromJSONTyped(json, ignoreDiscriminator) {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        'integrationSystem': !(0, runtime_1.exists)(json, 'integrationSystem') ? undefined : (0, IntegrationSystemType_1.IntegrationSystemTypeFromJSON)(json['integrationSystem']),
        'conversionCodeMappings': !(0, runtime_1.exists)(json, 'conversionCodeMappings') ? undefined : (json['conversionCodeMappings'].map(ConversionCodeMappingType_1.ConversionCodeMappingTypeFromJSON)),
        'links': !(0, runtime_1.exists)(json, 'links') ? undefined : (json['links'].map(InstanceLink_1.InstanceLinkFromJSON)),
        'warnings': !(0, runtime_1.exists)(json, 'warnings') ? undefined : (json['warnings'].map(WarningType_1.WarningTypeFromJSON)),
    };
}
exports.ConversionCodeMappingsToBeCreatedFromJSONTyped = ConversionCodeMappingsToBeCreatedFromJSONTyped;
function ConversionCodeMappingsToBeCreatedToJSON(value) {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        'integrationSystem': (0, IntegrationSystemType_1.IntegrationSystemTypeToJSON)(value.integrationSystem),
        'conversionCodeMappings': value.conversionCodeMappings === undefined ? undefined : (value.conversionCodeMappings.map(ConversionCodeMappingType_1.ConversionCodeMappingTypeToJSON)),
        'links': value.links === undefined ? undefined : (value.links.map(InstanceLink_1.InstanceLinkToJSON)),
        'warnings': value.warnings === undefined ? undefined : (value.warnings.map(WarningType_1.WarningTypeToJSON)),
    };
}
exports.ConversionCodeMappingsToBeCreatedToJSON = ConversionCodeMappingsToBeCreatedToJSON;
