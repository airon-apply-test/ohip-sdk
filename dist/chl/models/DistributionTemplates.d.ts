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
import type { DistributionTemplateType } from './DistributionTemplateType';
import type { HotelDistributionTemplatesType } from './HotelDistributionTemplatesType';
import type { InstanceLink } from './InstanceLink';
import type { WarningType } from './WarningType';
/**
 * Response object for fetching distribution templates based on search criteria.
 * @export
 * @interface DistributionTemplates
 */
export interface DistributionTemplates {
    /**
     * Information about a distribution template.
     * @type {Array<DistributionTemplateType>}
     * @memberof DistributionTemplates
     */
    distributionTemplates?: Array<DistributionTemplateType>;
    /**
     * Information about a hotel level distribution template.
     * @type {Array<HotelDistributionTemplatesType>}
     * @memberof DistributionTemplates
     */
    hotelsDistributionTemplates?: Array<HotelDistributionTemplatesType>;
    /**
     *
     * @type {Array<InstanceLink>}
     * @memberof DistributionTemplates
     */
    links?: Array<InstanceLink>;
    /**
     * Used in conjunction with the Success element to define a business error.
     * @type {Array<WarningType>}
     * @memberof DistributionTemplates
     */
    warnings?: Array<WarningType>;
}
/**
 * Check if a given object implements the DistributionTemplates interface.
 */
export declare function instanceOfDistributionTemplates(value: object): boolean;
export declare function DistributionTemplatesFromJSON(json: any): DistributionTemplates;
export declare function DistributionTemplatesFromJSONTyped(json: any, ignoreDiscriminator: boolean): DistributionTemplates;
export declare function DistributionTemplatesToJSON(value?: DistributionTemplates | null): any;
