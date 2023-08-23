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
import type { GuaranteeMappingType } from './GuaranteeMappingType';
import type { InstanceLink } from './InstanceLink';
import type { WarningType } from './WarningType';
/**
 * Request object for changing external system guarantees.
 * @export
 * @interface Guarantees
 */
export interface Guarantees {
    /**
     * Information about an external system guarantee mapping.
     * @type {Array<GuaranteeMappingType>}
     * @memberof Guarantees
     */
    guarantees?: Array<GuaranteeMappingType>;
    /**
     *
     * @type {Array<InstanceLink>}
     * @memberof Guarantees
     */
    links?: Array<InstanceLink>;
    /**
     * Used in conjunction with the Success element to define a business error.
     * @type {Array<WarningType>}
     * @memberof Guarantees
     */
    warnings?: Array<WarningType>;
}
/**
 * Check if a given object implements the Guarantees interface.
 */
export declare function instanceOfGuarantees(value: object): boolean;
export declare function GuaranteesFromJSON(json: any): Guarantees;
export declare function GuaranteesFromJSONTyped(json: any, ignoreDiscriminator: boolean): Guarantees;
export declare function GuaranteesToJSON(value?: Guarantees | null): any;
