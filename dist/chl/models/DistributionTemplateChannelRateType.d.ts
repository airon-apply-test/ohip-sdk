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
import type { SystemInfoType } from './SystemInfoType';
/**
 * Information about channel rates of a distribution template.
 * @export
 * @interface DistributionTemplateChannelRateType
 */
export interface DistributionTemplateChannelRateType {
    /**
     *
     * @type {SystemInfoType}
     * @memberof DistributionTemplateChannelRateType
     */
    systemInfo?: SystemInfoType;
    /**
     * Channel Rate Code where distribution template will be distributed to.
     * @type {string}
     * @memberof DistributionTemplateChannelRateType
     */
    channelRateCode?: string;
    /**
     * Channel Rate Category where distribution template will be distributed to.
     * @type {string}
     * @memberof DistributionTemplateChannelRateType
     */
    channelRateCategory?: string;
    /**
     * Channel Rate Level where distribution template will be distributed to.
     * @type {string}
     * @memberof DistributionTemplateChannelRateType
     */
    channelRateLevel?: string;
    /**
     * Regional Availability flag which will be the default value when rate code is distributed to channels.
     * @type {boolean}
     * @memberof DistributionTemplateChannelRateType
     */
    regionalAvailability?: boolean;
    /**
     * Rate update flag which will be the default value when rate code is distributed to channels.
     * @type {boolean}
     * @memberof DistributionTemplateChannelRateType
     */
    rateUpdate?: boolean;
    /**
     * Restriction update flag which will be the default value when rate code is distributed to channels.
     * @type {boolean}
     * @memberof DistributionTemplateChannelRateType
     */
    restrictionUpdate?: boolean;
    /**
     * Return SGA flag which will be the default value when rate code is distributed to channels.
     * @type {boolean}
     * @memberof DistributionTemplateChannelRateType
     */
    returnSGA?: boolean;
}
/**
 * Check if a given object implements the DistributionTemplateChannelRateType interface.
 */
export declare function instanceOfDistributionTemplateChannelRateType(value: object): boolean;
export declare function DistributionTemplateChannelRateTypeFromJSON(json: any): DistributionTemplateChannelRateType;
export declare function DistributionTemplateChannelRateTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): DistributionTemplateChannelRateType;
export declare function DistributionTemplateChannelRateTypeToJSON(value?: DistributionTemplateChannelRateType | null): any;
