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
import type { ChannelParameterMappingType } from './ChannelParameterMappingType';
import type { InstanceLink } from './InstanceLink';
import type { WarningType } from './WarningType';
/**
 * Response object for channel parameter(s) based on search criteria.
 * @export
 * @interface ChannelParameterMapping
 */
export interface ChannelParameterMapping {
    /**
     *
     * @type {ChannelParameterMappingType}
     * @memberof ChannelParameterMapping
     */
    channelParameterMapping?: ChannelParameterMappingType;
    /**
     *
     * @type {Array<InstanceLink>}
     * @memberof ChannelParameterMapping
     */
    links?: Array<InstanceLink>;
    /**
     * Used in conjunction with the Success element to define a business error.
     * @type {Array<WarningType>}
     * @memberof ChannelParameterMapping
     */
    warnings?: Array<WarningType>;
}
/**
 * Check if a given object implements the ChannelParameterMapping interface.
 */
export declare function instanceOfChannelParameterMapping(value: object): boolean;
export declare function ChannelParameterMappingFromJSON(json: any): ChannelParameterMapping;
export declare function ChannelParameterMappingFromJSONTyped(json: any, ignoreDiscriminator: boolean): ChannelParameterMapping;
export declare function ChannelParameterMappingToJSON(value?: ChannelParameterMapping | null): any;
