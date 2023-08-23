/**
 * OPERA Cloud Enterprise Configuration API
 * APIs to cater for Enterprise Configuration functionality in OPERA Cloud. <br /><br In this module, you can configure a variety of options related to your properties such as their locations, facilities, and local attractions. The available options are dependant on the active controls at your property.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 21.5.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
import type { InstanceLink } from './InstanceLink';
import type { TransportationType } from './TransportationType';
import type { WarningType } from './WarningType';
/**
 * Request object for creating transportation.
 * @export
 * @interface TransportationCriteria
 */
export interface TransportationCriteria {
    /**
     * Collection of hotel level transportation.
     * @type {Array<TransportationType>}
     * @memberof TransportationCriteria
     */
    transportationList?: Array<TransportationType>;
    /**
     *
     * @type {Array<InstanceLink>}
     * @memberof TransportationCriteria
     */
    links?: Array<InstanceLink>;
    /**
     * Used in conjunction with the Success element to define a business error.
     * @type {Array<WarningType>}
     * @memberof TransportationCriteria
     */
    warnings?: Array<WarningType>;
}
/**
 * Check if a given object implements the TransportationCriteria interface.
 */
export declare function instanceOfTransportationCriteria(value: object): boolean;
export declare function TransportationCriteriaFromJSON(json: any): TransportationCriteria;
export declare function TransportationCriteriaFromJSONTyped(json: any, ignoreDiscriminator: boolean): TransportationCriteria;
export declare function TransportationCriteriaToJSON(value?: TransportationCriteria | null): any;
