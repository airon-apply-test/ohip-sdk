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
import type { ChannelOrgMappingType } from './ChannelOrgMappingType';
import type { InstanceLink } from './InstanceLink';
import type { WarningType } from './WarningType';
/**
 * Request object for Create Channel to ORG and CRO Mapping.
 * @export
 * @interface ChannelOrgMapping
 */
export interface ChannelOrgMapping {
    /**
     *
     * @type {ChannelOrgMappingType}
     * @memberof ChannelOrgMapping
     */
    channelOrgMappingDetails?: ChannelOrgMappingType;
    /**
     *
     * @type {Array<InstanceLink>}
     * @memberof ChannelOrgMapping
     */
    links?: Array<InstanceLink>;
    /**
     * Used in conjunction with the Success element to define a business error.
     * @type {Array<WarningType>}
     * @memberof ChannelOrgMapping
     */
    warnings?: Array<WarningType>;
}
/**
 * Check if a given object implements the ChannelOrgMapping interface.
 */
export declare function instanceOfChannelOrgMapping(value: object): boolean;
export declare function ChannelOrgMappingFromJSON(json: any): ChannelOrgMapping;
export declare function ChannelOrgMappingFromJSONTyped(json: any, ignoreDiscriminator: boolean): ChannelOrgMapping;
export declare function ChannelOrgMappingToJSON(value?: ChannelOrgMapping | null): any;
