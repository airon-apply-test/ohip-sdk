"use strict";
/* tslint:disable */
/* eslint-disable */
/**
 * OPERA Cloud Channel Configuration API
 * APIs to cater for Channel Management functionality in OPERA Cloud. <br /><br /> Channel Management allows a property to configure and administer channels such as OTAs, and web channels, covering functionality such as channel configuration, availability, inventory and restrictions.<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 22.3.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
Object.defineProperty(exports, "__esModule", { value: true });
exports.ChannelGlobalDescriptionsToJSON = exports.ChannelGlobalDescriptionsFromJSONTyped = exports.ChannelGlobalDescriptionsFromJSON = exports.instanceOfChannelGlobalDescriptions = void 0;
const runtime_1 = require("../runtime");
const GenericGlobalDescriptionType_1 = require("./GenericGlobalDescriptionType");
const InstanceLink_1 = require("./InstanceLink");
const WarningType_1 = require("./WarningType");
/**
 * Check if a given object implements the ChannelGlobalDescriptions interface.
 */
function instanceOfChannelGlobalDescriptions(value) {
    let isInstance = true;
    return isInstance;
}
exports.instanceOfChannelGlobalDescriptions = instanceOfChannelGlobalDescriptions;
function ChannelGlobalDescriptionsFromJSON(json) {
    return ChannelGlobalDescriptionsFromJSONTyped(json, false);
}
exports.ChannelGlobalDescriptionsFromJSON = ChannelGlobalDescriptionsFromJSON;
function ChannelGlobalDescriptionsFromJSONTyped(json, ignoreDiscriminator) {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        'channelGlobalDescriptions': !(0, runtime_1.exists)(json, 'channelGlobalDescriptions') ? undefined : (json['channelGlobalDescriptions'].map(GenericGlobalDescriptionType_1.GenericGlobalDescriptionTypeFromJSON)),
        'links': !(0, runtime_1.exists)(json, 'links') ? undefined : (json['links'].map(InstanceLink_1.InstanceLinkFromJSON)),
        'warnings': !(0, runtime_1.exists)(json, 'warnings') ? undefined : (json['warnings'].map(WarningType_1.WarningTypeFromJSON)),
    };
}
exports.ChannelGlobalDescriptionsFromJSONTyped = ChannelGlobalDescriptionsFromJSONTyped;
function ChannelGlobalDescriptionsToJSON(value) {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        'channelGlobalDescriptions': value.channelGlobalDescriptions === undefined ? undefined : (value.channelGlobalDescriptions.map(GenericGlobalDescriptionType_1.GenericGlobalDescriptionTypeToJSON)),
        'links': value.links === undefined ? undefined : (value.links.map(InstanceLink_1.InstanceLinkToJSON)),
        'warnings': value.warnings === undefined ? undefined : (value.warnings.map(WarningType_1.WarningTypeToJSON)),
    };
}
exports.ChannelGlobalDescriptionsToJSON = ChannelGlobalDescriptionsToJSON;
