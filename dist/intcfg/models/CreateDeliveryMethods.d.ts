/**
 * OPERA Cloud Integration Configuration API
 * APIs catering to Integration Configuration in OPERA Cloud.  Operations such as get Hotel Interface Types, or get UDF mappings can be found in this module.<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 22.3.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
import type { ConfiguredExternalDeliveryMethodsType } from './ConfiguredExternalDeliveryMethodsType';
import type { InstanceLink } from './InstanceLink';
import type { WarningType } from './WarningType';
/**
 * Request to create the delivery method for a property/HUB/Chain and its configurations.
 * @export
 * @interface CreateDeliveryMethods
 */
export interface CreateDeliveryMethods {
    /**
     *
     * @type {ConfiguredExternalDeliveryMethodsType}
     * @memberof CreateDeliveryMethods
     */
    deliveryMethods?: ConfiguredExternalDeliveryMethodsType;
    /**
     *
     * @type {Array<InstanceLink>}
     * @memberof CreateDeliveryMethods
     */
    links?: Array<InstanceLink>;
    /**
     * Used in conjunction with the Success element to define a business error.
     * @type {Array<WarningType>}
     * @memberof CreateDeliveryMethods
     */
    warnings?: Array<WarningType>;
}
/**
 * Check if a given object implements the CreateDeliveryMethods interface.
 */
export declare function instanceOfCreateDeliveryMethods(value: object): boolean;
export declare function CreateDeliveryMethodsFromJSON(json: any): CreateDeliveryMethods;
export declare function CreateDeliveryMethodsFromJSONTyped(json: any, ignoreDiscriminator: boolean): CreateDeliveryMethods;
export declare function CreateDeliveryMethodsToJSON(value?: CreateDeliveryMethods | null): any;
