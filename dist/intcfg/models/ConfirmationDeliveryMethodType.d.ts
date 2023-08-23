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
import type { EmailDeliveryConfigurationType } from './EmailDeliveryConfigurationType';
import type { FaxDeliveryConfigurationType } from './FaxDeliveryConfigurationType';
import type { TextMessageDeliveryConfigurationType } from './TextMessageDeliveryConfigurationType';
/**
 * Confirmation letter delivery method configuration, settings for Text Message, Email and Fax Delivery of Confirmation Letters
 * @export
 * @interface ConfirmationDeliveryMethodType
 */
export interface ConfirmationDeliveryMethodType {
    /**
     *
     * @type {EmailDeliveryConfigurationType}
     * @memberof ConfirmationDeliveryMethodType
     */
    email?: EmailDeliveryConfigurationType;
    /**
     *
     * @type {FaxDeliveryConfigurationType}
     * @memberof ConfirmationDeliveryMethodType
     */
    fax?: FaxDeliveryConfigurationType;
    /**
     *
     * @type {TextMessageDeliveryConfigurationType}
     * @memberof ConfirmationDeliveryMethodType
     */
    textMessage?: TextMessageDeliveryConfigurationType;
}
/**
 * Check if a given object implements the ConfirmationDeliveryMethodType interface.
 */
export declare function instanceOfConfirmationDeliveryMethodType(value: object): boolean;
export declare function ConfirmationDeliveryMethodTypeFromJSON(json: any): ConfirmationDeliveryMethodType;
export declare function ConfirmationDeliveryMethodTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): ConfirmationDeliveryMethodType;
export declare function ConfirmationDeliveryMethodTypeToJSON(value?: ConfirmationDeliveryMethodType | null): any;
